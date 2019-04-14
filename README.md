# Mycelia
_Stealthy distributed communications network_

### Summary
This seemed like a good idea at the time.

### Innards
Mycelia binaries uniquely mutate upon execution to avoid static detection, the structure of the built application is as follows:

```
| Common application headers
| Application
| Set of all known hosts
```

When run, the server calls out to hosts in the set until it is able to resolve one. If there are no hosts, it starts listening. If it cannot resolve any hosts, the program exits.

Once it resolves a host (or is no host list), it starts its own dns responder. This responder will respond to requests in the same way a standard dns server would, but for specially crafted packets will reply with its host list or other functionality.

**Standard DNS Packet (RFC 1035)**
```
 0   1   2   3   4   5   6   7   8   9   10  11  12  13  14  15
 [ ID                                                         ]
 [QR ][OPCODE    ][AA][TC][RD][RA][Z          ][RCODE         ]
 [ QDCOUNT                                                    ]
 [ ADCOUNT                                                    ]
 [ NSCOUNT                                                    ]
 [ ARCOUNT                                                    ]
```

**Mycelia DNS Packet (RFC -0x0)**
```
 0   1   2   3   4   5   6   7   8   9   10  11  12  13  14  15
 [ junk that looks like corrupt udp traffic                   ]
 [QR ][XORSTRING ][AA][TC][RD][RA][Z          ][RCODE         ]
 +------------------------------------------------------------+
 |xor(payload,(xorstring,clienthost))                         |
 |                                                            |
 |                                                            |
 +------------------------------------------------------------+
```

The daemon utilizes golang's dns library. Upon receiving a query, it attempts to process it as if it were a legitimate request.
Processing it successfully will yield a legitimate dns response. Upon failure, mycelia will attempt to decode the packet data.

