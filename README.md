# goms
CN MapleStory Server Demo.  
Only implemented the login server.  
The channel server need much time to overcome... 
## Runtime
- MapleVersion: CMS138.1
- Golang version at least 1.21 to support slog
- Database: mongodb latest version
- [Convert a Standalone mongod to a Replica Set](https://www.mongodb.com/docs/manual/tutorial/convert-standalone-to-replica-set/)
- Convert wz files to nx files and put them in nxfile directory
    - suggest use [Harepacker-resurrected](https://github.com/lastbattle/Harepacker-resurrected)
- Must forward game target ip 221.231.130.70 to server ip
    - suggest use [goms-tool](https://github.com/zhyonc/goms-tool) launcher to patch game
## References
- [KMS1029IDB](https://forum.ragezone.com/threads/leak-maplestory-korea-test-v-1029-raw-exe-w-debuggables.1100141/)
- [GMS178Server](https://forum.ragezone.com/threads/java-v178-swordie-source-named-v178-1-idb.1220875/)
- [MapleSharkPacket](https://github.com/zh3305/MapleShark---Scripts)
## Packet
|Header|AESOFB|Note|
|:---:|:---:|:---:|
|4 Bytes|Any Bytes|Except for the first packet|
### Decode
PacketLen = (Header[0]+Header[1]*0x100) ^ (Header[2]+Header[3]*0x100)
### Encode
- sVersion = (^clientVer >> 8 & 0xFF) | ((^clientVer << 8) & 0xFF00)
- a = int(sendIV[3])
- a |= int(sendIV[2])<<8
- a ^= sVersion
- b = ((PacketLen << 8) & 0xFF00) | (PacketLen >> 8)
- c = a ^ b
- Header = [a>>8, a, c>>8, b]
### PacketFormat
|Opcode|Data|Note|
|:---:|:---:|:---:|
|2 Bytes|Any Bytes|Except for the first packet|
### The First Packet
|Name|PacketLen|Version|MinorVersionLen|MinorVersion|RecvIV|SendIV|Region|IsTestServer|Note|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
|Connect|2 Bytes|2 Bytes|2 Bytes|1 Byte|4 Bytes|4 Bytes|1 Bytes|1 Byte|The first packet|