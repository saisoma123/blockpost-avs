// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "../src/BlockPostServiceManager.sol";

contract BlockPostServiceManagerTestProxy is BlockPostServiceManager {
    constructor(
        IAVSDirectory _avsDirectory,
        IRegistryCoordinator _registryCoordinator,
        IStakeRegistry _stakeRegistry
    )
        BlockPostServiceManager(
            _avsDirectory,
            _registryCoordinator,
            _stakeRegistry
        )
    {}

    // Expose internal getter functions for testing
    function getMessage(
        uint256 _messageId
    ) public view returns (string memory) {
        return _getMessage(_messageId);
    }

    function isMessageValidated(uint256 _messageId) public view returns (bool) {
        return _isMessageValidated(_messageId);
    }

    function getAddressForMessage(
        uint256 _messageId
    ) public view returns (address) {
        return _getAddressForMessage(_messageId);
    }
}
