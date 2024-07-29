// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "@eigenlayer/contracts/libraries/BytesLib.sol";
import "@openzeppelin-upgrades/contracts/utils/cryptography/ECDSAUpgradeable.sol";
import "@eigenlayer/contracts/core/DelegationManager.sol";
import "eigenlayer-middleware/src/ServiceManagerBase.sol";
import "eigenlayer-middleware/src/unaudited/ECDSAStakeRegistry.sol";
import "@openzeppelin-upgrades/contracts/utils/cryptography/ECDSAUpgradeable.sol";
import "@eigenlayer/contracts/permissions/Pausable.sol";
import {IRegistryCoordinator} from "eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";

contract BlockPostServiceManager is ServiceManagerBase, Pausable {
    uint256 private messageId;

    mapping(uint256 => string) private messages; // Stores all messages
    mapping(uint256 => bool) private messageValidated; // Keeps track of message validation
    mapping(uint256 => address) private idstoAddress; // Maps message ids to addresses for retrieval checks

    event MessageSubmitted(uint256 indexed messageId, string message); // Emits event to operator listener
    event MessageValidated(
        uint256 indexed messageId,
        string message,
        address sender
    ); // Emits provable on chain event to show message was validated properly

    // Constructs with ServiceManager to keep with AVS requirements
    constructor(
        IAVSDirectory _avsDirectory,
        IRegistryCoordinator _registryCoordinator,
        IStakeRegistry _stakeRegistry
    ) ServiceManagerBase(_avsDirectory, _registryCoordinator, _stakeRegistry) {
        messageId = 0;
    }

    // Users use this method to submit a message that is tracked by the operator
    function submitMessage(
        string memory _message
    ) public whenNotPaused returns (uint256) {
        require(bytes(_message).length > 0, "Message cannot be empty");

        uint256 currMessageId = messageId; // Saves curr message id value
        messageId++; // Increments to new value

        idstoAddress[currMessageId] = msg.sender; // Maps the curr value to user address

        emit MessageSubmitted(currMessageId, _message); // Emits event to operator
        return currMessageId;
    }

    // The operator calls this method to send the signature back to ServiceManager for recovery and storage
    function storeValidatedMessage(
        uint256 _messageId,
        string memory _message,
        bytes memory _signature
    ) public whenNotPaused {
        // Checks if message was validated already, avoids re-entry
        require(!messageValidated[_messageId], "Message already validated");

        // Verify the signature
        bytes32 messageHash = keccak256(abi.encodePacked(_message));
        bytes32 ethSignedMessageHash = ECDSAUpgradeable.toEthSignedMessageHash(
            messageHash
        );
        address signer = ECDSAUpgradeable.recover(
            ethSignedMessageHash,
            _signature
        );

        // Checks if operator was the one who sent the signature
        require(signer == msg.sender, "Message signer is not operator");

        // Stores validated message
        messages[_messageId] = _message;
        messageValidated[_messageId] = true;

        emit MessageValidated(_messageId, _message, idstoAddress[_messageId]);
    }

    // Users use this to retrieve a message only they stored
    function retrieveMessage(
        uint256 _messageId
    ) public view returns (string memory) {
        // Checks for message length, validation, and if user retrieving is user who sent before
        require(
            bytes(messages[_messageId]).length > 0,
            "Message ID does not exist"
        );
        require(messageValidated[_messageId], "Message not validated");
        require(
            idstoAddress[_messageId] == msg.sender,
            "This is not your message!"
        );
        return messages[_messageId];
    }

    // These functions are for accessing mapping info for testers
    function _getMessage(
        uint256 _messageId
    ) internal view returns (string memory) {
        return messages[_messageId];
    }

    function _isMessageValidated(
        uint256 _messageId
    ) internal view returns (bool) {
        return messageValidated[_messageId];
    }

    function _getAddressForMessage(
        uint256 _messageId
    ) internal view returns (address) {
        return idstoAddress[_messageId];
    }
}
