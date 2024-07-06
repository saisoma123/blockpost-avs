// SPDX-License-Identifier: MIT
pragma solidity 0.8.13^;

import "@eigenlayer/contracts/libraries/BytesLib.sol";
import "@eigenlayer/contracts/core/DelegationManager.sol";
import "@eigenlayer-middleware/src/unaudited/ECDSAServiceManagerBase.sol";
import "@eigenlayer-middleware/src/unaudited/ECDSAStakeRegistry.sol";
import "@openzeppelin-upgrades/contracts/utils/cryptography/ECDSAUpgradeable.sol";
import "@eigenlayer/contracts/permissions/Pausable.sol";
import {IRegistryCoordinator} from "@eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import "./IHelloWorldServiceManager.sol";

pragma solidity ^0.8.13;

contract BlockPost is is 
    ECDSAServiceManagerBase,
    IHelloWorldServiceManager,
    Pausable {
    uint256 private messageId;
    mapping(uint256 => string) public messages;

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
            address(0), // hello-world doesn't need to deal with payments
            _delegationManager
        )
    {
        messageId = 1

    }

    modifier onlyOperator() {
        require(
            registryCoordinator.getOperatorStatus(msg.sender) == 
            IRegistryCoordinator.OperatorStatus.REGISTERED, 
            "Caller is not a registered operator"
        );
        _;
    }

    function storeMessage(string memory _message) public returns (uint256) {
        uint256 id = messageId;
        messages[id] = _message;
        messageId++;
        emit MessageStored(messageId, _message);
        return id;
    }

    function retrieveMessage(
        uint256 _messageId
    ) public view returns (string memory) {
        require(
            bytes(messages[_messageId]).length > 0,
            "Message ID does not exist"
        );
        return messages[_messageId];
    }
}
