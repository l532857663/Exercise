#!/usr/bin/env python3
#! -*- coding:utf-8 -*-

import hashlib

def get_sha1prng_key(key):
    signature = hashlib.sha1(key.encode()).digest()
    signature = hashlib.sha1(signature).digest()
    return ''.join(['%02x' % i for i in signature]).upper()[:32]

if __name__ == "__main__":
    print("Start")
    key = "1234567890123456"
    print(key)
    key32 = get_sha1prng_key(key)
    print(key32)
    print("End")
