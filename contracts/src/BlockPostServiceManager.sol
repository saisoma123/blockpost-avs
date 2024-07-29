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

    mapping(uint256 => string) private messages;
    mapping(uint256 => bool) private messageValidated;
    mapping(uint256 => address) private idstoAddress;

    event MessageSubmitted(uint256 indexed messageId, string message);
    event MessageValidated(
        uint256 indexed messageId,
        string message,
        address sender
    );

    constructor(
        IAVSDirectory _avsDirectory,
        IRegistryCoordinator _registryCoordinator,
        IStakeRegistry _stakeRegistry
    ) ServiceManagerBase(_avsDirectory, _registryCoordinator, _stakeRegistry) {
        messageId = 0;
    }

    function submitMessage(
        string memory _message
    ) public whenNotPaused returns (uint256) {
        require(bytes(_message).length > 0, "Message cannot be empty");

        uint256 currMessageId = messageId;
        messageId++;

        idstoAddress[currMessageId] = msg.sender;

        emit MessageSubmitted(currMessageId, _message);
        return currMessageId;
    }

    function storeValidatedMessage(
        uint256 _messageId,
        string memory _message,
        bytes memory _signature
    ) public whenNotPaused {
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

        require(signer == msg.sender, "Message signer is not operator");

        messages[_messageId] = _message;
        messageValidated[_messageId] = true;

        emit MessageValidated(_messageId, _message, idstoAddress[_messageId]);
    }

    function retrieveMessage(
        uint256 _messageId
    ) public view returns (string memory) {
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
