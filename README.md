# WakeUp

Simple utility to wake computers on LAN.

Exposes a function, so can be used as a lib.

Has no external dependencies.

## Usage

```sh
wakeup 192.168.0.255 DE:AD:BE:EF:CA:FE
```

Where `"192.168.0.255"` is a broadcast address for your LAN subnetwork, and `"DE:AD:BE:EF:CA:FE"` is a MAC-address of the computer that you want to wake up.
Only IPv4 is supported.

## How it works

[Wikipedia](https://en.wikipedia.org/wiki/Wake-on-LAN)

Basically, to wake any computer on LAN that has this feature enabled in BIOS, all you need to do is to send a special "magic" network packet:

> The magic packet is a broadcast frame containing anywhere within its payload 6 bytes of all 255 (FF FF FF FF FF FF in hexadecimal), followed by sixteen repetitions of the target computer's 48-bit MAC address, for a total of 102 bytes.

## License

MIT
