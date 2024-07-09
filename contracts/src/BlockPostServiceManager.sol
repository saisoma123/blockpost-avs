// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "eigenlayer/contracts/libraries/BytesLib.sol";
import "eigenlayer/contracts/core/DelegationManager.sol";
import "eigenlayer-middleware/src/ServiceManagerBase.sol";
import "eigenlayer-middleware/src/unaudited/ECDSAStakeRegistry.sol";
import "openzeppelin-upgrades/contracts/utils/cryptography/ECDSAUpgradeable.sol";
import "eigenlayer/contracts/permissions/Pausable.sol";
import {IRegistryCoordinator} from "eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import "./IBlockPostServiceManager.sol";

pragma solidity ^0.8.13;

contract BlockPostServiceManager is
    ServiceManagerBase,
    IHelloWorldServiceManager,
    Pausable
{
    uint256 private messageId;

    mapping(uint256 => string) public messages;
    // mapping(uint256 => bool) public messageValidated;
    mapping(uint256 => bytes) public messageSignatures;

    event MessageStored(uint256 indexed messageId, string message);
    event MessageSubmitted(uint256 indexed messageId, string message);
    event MessageValidated(uint256 indexed messageId, string message);

    constructor(
        address _avsDirectory,
        address _stakeRegistry,
        address _delegationManager
    )
        ECDSAServiceManagerBase(
            _avsDirectory,
            _stakeRegistry,
            address(0),
            _delegationManager
        )
    {
        messageId = 1;
    }

    modifier onlyOperator() {
        require(
            registryCoordinator.getOperatorStatus(msg.sender) ==
                IRegistryCoordinator.OperatorStatus.REGISTERED,
            "Caller is not a registered operator"
        );
        _;
    }

    function submitMessage(
        string memory _message
    ) public whenNotPaused returns (uint256) {
        require(bytes(_message).length > 0, "Message cannot be empty");

        uint256 messageId = latestMessageId;
        latestMessageId++;

        emit MessageSubmitted(messageId, _message);
        return messageId;
    }

    function storeValidatedMessage(
        uint256 _messageId,
        string memory _message,
        bytes memory _signature
    ) public onlyOperator whenNotPaused {
        //require(bytes(_message).length > 0, "Message cannot be empty");
        require(!messageValidated[_messageId], "Message already validated");

        // Verify the signature
        bytes32 messageHash = keccak256(abi.encodePacked(_message));
        bytes32 ethSignedMessageHash = messageHash.toEthSignedMessageHash();
        address signer = ethSignedMessageHash.recover(_signature);

        require(operators[signer], "Invalid signature");

        messages[_messageId] = _message;
        // messageValidated[_messageId] = true;
        messageSignatures[_messageId] = _signature;

        emit MessageValidated(_messageId, _message, _signature);
    }

    /** 
    function storeMessage(string memory _message) public returns (uint256) {
        uint256 id = messageId;
        messages[id] = _message;
        messageId++;
        emit MessageStored(messageId, _message);
        return id;
    }
    */

    function retrieveMessage(
        uint256 _messageId
    ) public view returns (string memory) {
        require(
            bytes(messages[_messageId]).length > 0,
            "Message ID does not exist"
        );
        require(messageValidated[_messageId], "Message not validated");
        return messages[_messageId];
    }
}
