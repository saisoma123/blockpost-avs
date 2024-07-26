// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "../src/BlockPostServiceManager.sol";
import "@openzeppelin-upgrades/contracts/utils/cryptography/ECDSAUpgradeable.sol";
import "eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";

contract BlockPostServiceManagerTest is Test {
    BlockPostServiceManager public blockPostServiceManager;
    address public owner;
    address public user;
    address public validator;
    uint256 validatorKey;
    uint256 public messageId;
    bytes32 public messageHash;
    bytes public signature;

    function setUp() public {
        owner = address(0x1);
        user = address(0x2);

        validatorKey = uint256(keccak256(abi.encodePacked("validator")));
        validator = vm.addr(validatorKey);

        IAVSDirectory avsDirectory = IAVSDirectory(address(0x4));
        IRegistryCoordinator registryCoordinator = IRegistryCoordinator(
            address(0x5)
        );
        IStakeRegistry stakeRegistry = IStakeRegistry(address(0x6));

        blockPostServiceManager = new BlockPostServiceManager(
            avsDirectory,
            registryCoordinator,
            stakeRegistry
        );

        vm.startPrank(owner);
    }

    function testSubmitMessage() public {
        vm.stopPrank();
        vm.startPrank(user);
        messageId = blockPostServiceManager.submitMessage("Hello, World!");

        address validated = blockPostServiceManager.idstoAddress(messageId);
        assertEq(validated, user);
    }

    function testStoreValidatedMessage() public {
        vm.stopPrank();
        vm.startPrank(user);
        messageId = blockPostServiceManager.submitMessage("Hello, World!");
        vm.stopPrank();

        vm.startPrank(validator);
        // Validator signs the message
        (messageHash, signature) = signMessage("Hello, World!");

        // Ensure validator is the msg.sender when calling storeValidatedMessage
        blockPostServiceManager.storeValidatedMessage(
            messageId,
            "Hello, World!",
            signature
        );
        vm.stopPrank();

        string memory message = blockPostServiceManager.messages(messageId);
        bool validated = blockPostServiceManager.messageValidated(messageId);
        assertEq(message, "Hello, World!");
        assertEq(validated, true);
    }

    function testRetrieveMessage() public {
        vm.stopPrank();
        vm.startPrank(user);
        messageId = blockPostServiceManager.submitMessage("Hello, World!");

        vm.stopPrank();
        vm.startPrank(validator);
        (messageHash, signature) = signMessage("Hello, World!");

        blockPostServiceManager.storeValidatedMessage(
            messageId,
            "Hello, World!",
            signature
        );

        vm.stopPrank();
        vm.startPrank(user);
        string memory retrievedMessage = blockPostServiceManager
            .retrieveMessage(messageId);
        assertEq(retrievedMessage, "Hello, World!");
    }

    function testSubmitEmptyMessage() public {
        vm.stopPrank();
        vm.startPrank(user);
        vm.expectRevert("Message cannot be empty");
        blockPostServiceManager.submitMessage("");
    }

    function testStoreValidatedMessageTwice() public {
        vm.stopPrank();
        vm.prank(user);
        messageId = blockPostServiceManager.submitMessage("Hello, World!");
        vm.stopPrank();

        vm.prank(validator);
        (messageHash, signature) = signMessage("Hello, World!");

        blockPostServiceManager.storeValidatedMessage(
            messageId,
            "Hello, World!",
            signature
        );

        vm.expectRevert("Message already validated");
        blockPostServiceManager.storeValidatedMessage(
            messageId,
            "Hello, World!",
            signature
        );
    }

    function testRetrieveMessageNotStored() public {
        vm.stopPrank();
        vm.startPrank(user);
        messageId = blockPostServiceManager.submitMessage("Hello, World!");
        vm.stopPrank();

        vm.startPrank(user);
        vm.expectRevert("Message ID does not exist");
        blockPostServiceManager.retrieveMessage(messageId);
    }

    function testRetrieveMessageNotOwner() public {
        vm.stopPrank();
        vm.prank(user);
        messageId = blockPostServiceManager.submitMessage("Hello, World!");
        vm.stopPrank();

        vm.prank(validator);
        (messageHash, signature) = signMessage("Hello, World!");

        blockPostServiceManager.storeValidatedMessage(
            messageId,
            "Hello, World!",
            signature
        );
        vm.stopPrank();

        vm.expectRevert("This is not your message!");
        vm.prank(owner);
        blockPostServiceManager.retrieveMessage(messageId);
    }

    function testRetrieveMessageWithInvalidSignature() public {
        vm.stopPrank();
        vm.startPrank(user);
        messageId = blockPostServiceManager.submitMessage("Hello, World!");
        vm.stopPrank();

        vm.startPrank(validator);
        (messageHash, signature) = signMessage("Invalid Message");
        vm.expectRevert("Message signer is not operator");
        blockPostServiceManager.storeValidatedMessage(
            messageId,
            "Hello, World!",
            signature
        );
    }

    function testRetrieveMultipleMessages() public {
        vm.stopPrank();
        vm.startPrank(user);

        uint256[] memory messageIds = new uint256[](15);
        string[15] memory messages = [
            "Message 1",
            "Message 2",
            "Message 3",
            "Message 4",
            "Message 5",
            "Message 6",
            "Message 7",
            "Message 8",
            "Message 9",
            "Message 10",
            "Message 11",
            "Message 12",
            "Message 13",
            "Message 14",
            "Message 15"
        ];

        // Submit 15 messages
        for (uint i = 0; i < 15; i++) {
            messageIds[i] = blockPostServiceManager.submitMessage(messages[i]);
        }
        vm.stopPrank();

        vm.startPrank(validator);
        // Validate 15 messages
        for (uint i = 0; i < 15; i++) {
            (messageHash, signature) = signMessage(messages[i]);
            blockPostServiceManager.storeValidatedMessage(
                messageIds[i],
                messages[i],
                signature
            );
        }
        vm.stopPrank();

        vm.startPrank(user);
        // Retrieve and check 15 messages
        for (uint i = 0; i < 15; i++) {
            string memory retrievedMessage = blockPostServiceManager
                .retrieveMessage(messageIds[i]);
            assertEq(retrievedMessage, messages[i]);
        }
        vm.stopPrank();
    }

    function signMessage(
        string memory _message
    ) internal view returns (bytes32, bytes memory) {
        bytes32 hash = keccak256(abi.encodePacked(_message));
        bytes32 ethSignedMessageHash = ECDSAUpgradeable.toEthSignedMessageHash(
            hash
        );

        (uint8 v, bytes32 r, bytes32 s) = vm.sign(
            validatorKey,
            ethSignedMessageHash
        );

        bytes memory sig = abi.encodePacked(r, s, v);
        return (ethSignedMessageHash, sig);
    }
}
