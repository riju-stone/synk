# Synk
Simple self-hostable service to sync clipboard and files across all your devices

## Architecture
- API
  - Utliizes MultiCast DNS to detect LAN connected devices (first time setup)
  - Provides necessary endpoints to query STUN servers
  - Generates and returns connection SDPs
- The SDPs returned from the API is used to form a stable & secure WebRTC conection
- Data transfer is initiated


## References
- [always online stun](https://github.com/pradt2/always-online-stun)
- [sendfa](https://github.com/0xLaurens/sendfa.st)