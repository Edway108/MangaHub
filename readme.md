MangaHub Testing Guide

1. Run the Application

Open terminal and go to the cmd/dev directory:

cd cmd/dev
go run .

==================================================
UC-001: User Registration
==================================================

Register a new account:

mangahub register <username> <password>

Example:
mangahub register john 123456

==================================================
UC-002: User Authentication
==================================================

Login with an existing account:

mangahub login <username> <password>

Example:
mangahub login john 123456

==================================================
UC-003: Search Manga
==================================================

Search manga by keyword:

mangahub manga search <keyword>

Example:
mangahub manga search one

==================================================
UC-004: View Manga Details
==================================================

Retrieve manga details using gRPC:

mangahub grpc get <manga-id>

Example:
mangahub grpc get m1

Expected Output:
One Piece - Oda Eiichiro - Pirate adventure

==================================================
UC-005: Add Manga to Library
==================================================

Add manga to library:

mangahub library add <manga-id>

Example:
mangahub library add m1

==================================================
UC-006: Update Reading Progress
==================================================

Update reading progress:

mangahub progress update --chapter <chapter-number> --manga-id <manga-id>

Example:
mangahub progress update --chapter 32 --manga-id m40

Expected Output:
✓ Progress updated successfully

==================================================
UC-007: Connect to TCP Sync Server
==================================================

Connect to TCP sync server:

mangahub sync connect

==================================================
UC-009: Register for UDP Notifications
==================================================

Subscribe for UDP notifications:

mangahub subscribe

This command will open a listening channel for notifications.

==================================================
UC-010: Send Chapter Release Notification
==================================================

Send update notification:

mangahub progress update --chapter 32 --manga-id m40

Expected Output:
✓ Progress updated successfully

Other subscribed clients should receive the notification automatically.

==================================================
UC-011: Join Chat
==================================================

Join chat room:

mangahub chat

Open 2 terminals and login with provided accounts:

Terminal 1:
mangahub login quantran 123456

Terminal 2:
mangahub login tanminh 123456

Then run:
mangahub chat

==================================================
UC-012: Send Chat Message
==================================================

Type any message and press Enter.

Example:
Hello everyone

Use:
/help

to see more commands.

==================================================
UC-013: Handle User Disconnection
==================================================

Leave chat safely:

/quit

==================================================
UC-014: Retrieve Manga via gRPC
==================================================

Retrieve manga information through gRPC:

mangahub grpc get m1

Expected Output:
One Piece - Oda Eiichiro - Pirate adventure

==================================================
UC-015: Search Manga via gRPC
==================================================

Search manga via gRPC:

mangahub manga search one

==================================================
UC-016: Update Progress via gRPC
==================================================

Update reading progress through gRPC:

mangahub grpc progress update --chapter 109 --manga-id m12

Expected Output:
✓ Progress updated successfully

==================================================
NOTES
==================================================

- Make sure the server is running before testing client features.
- Some features require multiple terminals:
  + Chat
  + UDP notifications
  + TCP synchronization
- Use /help inside chat mode for additional commands.
