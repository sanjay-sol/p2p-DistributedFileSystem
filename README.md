# Peer-to-Peer File Storage and Retrieval System

## Table of Contents
- [Peer-to-Peer File Storage and Retrieval System](#peer-to-peer-file-storage-and-retrieval-system)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Project Overview](#project-overview)
  - [Architecture](#architecture)
  - [Flowchart](#flowchart)
  - [Components](#components)
    - [Client Nodes](#client-nodes)
    - [Peer Nodes](#peer-nodes)
    - [Directory Server](#directory-server)
  - [Definitions](#definitions)
  - [Future Enhancements](#future-enhancements)
  - [Conclusion](#conclusion)

## Introduction
This project aims to develop a decentralized peer-to-peer (P2P) file storage and retrieval system using Go. The focus is on distributed design, secure file handling, and efficient data retrieval, ensuring data integrity and availability.

## Project Overview
The system leverages P2P networking principles to distribute file storage across multiple nodes, ensuring redundancy and security. Users can store and retrieve files in a decentralized manner, eliminating the need for a central server.

## Architecture
![Architecture Diagram](/p2p_cloud.png)

The architecture consists of the following components:
- **Client Nodes**: Nodes operated by end-users that interact with the network to upload and download files.
- **Peer Nodes**: Nodes that store and manage files, ensuring data redundancy and availability.
- **Directory Server**: A lightweight server that maintains a directory of nodes and file locations, facilitating efficient file retrieval.

## Flowchart
![Flowchart Diagram](/p2p.png)

The flow of operations in the system:
1. **File Upload**:
   - The client node splits the file into chunks.
   - Each chunk is hashed and encrypted.
   - Chunks are distributed to multiple peer nodes.
   - The directory server is updated with the file metadata and chunk locations.
   
2. **File Retrieval**:
   - The client requests the file from the directory server.
   - The directory server provides the locations of the file chunks.
   - The client retrieves and decrypts the chunks from the peer nodes.
   - The chunks are reassembled into the original file.

## Components
### Client Nodes
- Responsible for file uploads and downloads.
- Interacts with the directory server and peer nodes.

### Peer Nodes
- Stores file chunks.
- Ensures data redundancy through replication.
- Participates in the P2P network to maintain data availability.

### Directory Server
- Maintains a mapping of file metadata to peer nodes.
- Facilitates efficient file retrieval by providing chunk locations.

## Definitions
- **Peer-to-Peer (P2P) Network**: A decentralized network where each participant (node) acts as both a client and a server.
- **Chunking**: The process of dividing a file into smaller pieces for distribution and storage.
- **Hashing**: Creating a unique identifier for each file chunk to ensure data integrity.
- **Encryption**: Securing file chunks to prevent unauthorized access.

## Future Enhancements
- Implementing advanced encryption techniques for enhanced security.
- Developing a more robust directory server with failover capabilities.
- Integrating blockchain technology for immutable file metadata storage.
- Implementing an incentive mechanism for peer node participation.

## Conclusion
This P2P file storage and retrieval system aims to provide a decentralized solution for secure and efficient data storage. By leveraging P2P principles, the system ensures data redundancy, security, and availability without relying on a central authority.

---

Feel free to contribute to this project by suggesting improvements, reporting issues, or submitting pull requests. Let's build a more decentralized future together!
