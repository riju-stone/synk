# Synk
Simple self-hostable service to sync clipboard and files across all your devices

## Architecture
- API
  - Utliizes MultiCast DNS to detect LAN connected devices (first time setup)
  - Provides necessary endpoints to query STUN servers
  - Generates and returns connection SDPs
- The SDPs returned from the API is used to form a stable & secure WebRTC conection
- Data transfer is initiated

## STUN Server List (Working as of 12th May 2025)

- stun:stun.l.google.com:19302
- stun:stun.l.google.com:5349
- stun:stun1.l.google.com:3478
- stun:stun1.l.google.com:5349
- stun:stun2.l.google.com:19302
- stun:stun2.l.google.com:5349
- stun:stun3.l.google.com:3478
- stun:stun3.l.google.com:5349
- stun:stun4.l.google.com:19302
- stun:stun4.l.google.com:5349

## Self-Hosted Stun Server


## References
- [always online stun](https://github.com/pradt2/always-online-stun)
- [sendfa](https://github.com/0xLaurens/sendfa.st)