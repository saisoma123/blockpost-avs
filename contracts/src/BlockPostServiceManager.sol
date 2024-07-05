// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

contract BlockPostServiceManager {
    uint256 private messageId;
    mapping(uint256 => string) public messages;

    event MessageStored(uint256 indexed messageId, string message);

    constructor() {
        messageId = 1;
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
