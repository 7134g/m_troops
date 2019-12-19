# 连续相同的字符串片段比对

a = "S6em9W2e5)F(QwT4nZDeJ(CPY8f)Eb)zQXI91bjhnkG7n46EmlG3MtU4(tEzqSjb)p4BJvJZFtj1Dawdiuq65t3FZTO6xgcAIwRXAoezAzXjOSJJhxKAOCzNrGIUpt2WzFFUHrae64yatzlx77hfIrk)WaziUr0TY2LNFov5cQJ8TCWdqWyNfAL4aUJqqehpXGsQH()1QWNzprXVAuDdb4I6cOjbf0bxm)TZ7A4lKNZ09c8RnfCMoJ6lPmMQG84xrt46hgjBDCjm)jUaykIZF9jRe1qzOHN9DWVTVtNMKTn9q6fVoITg6Uhbje9LHF0uMwZBLusZ1I3GZv3pl7flsewzKobHq2XyWEys)Ea2wx1YDWlCTEFwb3e6votnq9dk)BbAvMgEsz1kK3lVvu8wHlYE(9sY9XmgINWIFsSyWfXccScGD2q7rbViSXMPoVKDcjNgcDUAuejaXNg(izTxA2WPI0bOuFGwqySNjqzYMFnwbcmcOnzHQMGfWUFnkVxoh8NpR67F0s0AlLns6sby1axHgxGI6wVBl9y15cVZkPVba1KI94bhWq7sXmnLuN3rqcdgGBI4RgKMcmFl5skq9rpMe(pUkArX(M7yWuyYMjDDyd7t(Bj1X6Q57BregWMgRKU4T1RoCrghZPO(YEgsS8BC58gomAovKd6JEYqSeb9gG2pVgtSQgef)dlXQZAyM6UZ74OYM9oaM3lk1Gmrg1DaoT7VBrXzWhqdEPXQoxikauTFx9cVV9Lmwaw9XvL04QrO6AuDbpv0hO0LXQHRtf0qBO0SthMEV3PEUgP21dWfpA559OLVK2kKLKRQDBogGdQPRU0pvVvkpVZuDeSJVmA..96ee6eaab311e178905d07d6af705dcb21b1940275757ba22a2d181bbeaa82dbdb40830fd9bb89e8974b75af4aa037ecce411ddbe2862a9d95225727382e25205cf09f319813f080d96b59d69a2aa869e9313754bbeb542c8896fcd3100a8904857c8a9fb14befc29119b316aba57e4f7481840ebe11a7b48fb80f0c7def516b"
b = "S6em9W2e5)F(QwT4nZDeJ(CPY8f)Eb)zQXI91bjhnkE0djiXsVgMUlh08m4XWzFVRdzB8povpIZYLH7fb1)feCZ8G88wJDLsC)(JgvOtIkvYpIVIgI11kfd2dX0YnZh8Wou(GQUnjm13nVErQZEhfrBr0hvFv0PVU5aSBuwNZH5nvrYXUkw(4vN1t9z8LaPW(hvuBwkSK7vhifbYSGaFJxoZUJwOGlAztJe93o7GJD8QIAGduRS(H9S6v3kVIlHJWhNc1YAZxSGMpggVmIMnLjubJMJVVVIu8OPGfX1538gu8aICkCu(wtCw7jL(cd2tuLKKg(kSZ79l61iqBBhYlIsK4seHyGnv31yp(ScZpYr5S4RIDwecMlHO6TTW(Hm2amOjzJ4ZUpN2muyj(WfgxkmoTNpjKclRMbrS(sjGJ0xqHPY)u9p74dK0jEoBoigjfTZ)TkS9201DkSQ9g8xrxZDzaNb(m8X7oac)z9FqDNFkHW0ivibdeyeUrAotrNcfLiV7b2TWJz5TSmgVD1itvdvZOUJxJtA61YSnV2HqLYnUBLEzHSkE130A98WWIwU3FmqrOH)A3TcUy(PSEPBIsz0cYW4WkqD(OT1Xc3SCm4yReseSTzfSJAGL2TN9)yr2CZQMWVhkLlbxn2rHhV728AEqkP6XzNfXqyXz1Bd)8I9g36lJ8yDwqY1Wu0sRcErxSxztjk7BcDlIlseALZURys9SfRjjmk2BWvezxH1t4zXo1aBRmWVU6VkygM8QQ1GhDXMYxAavv(jTSE6aIQ)s)g..9c0ff9ce9c1e721c8f8f157f7ef9712a382e86e560d2d6a03005d1a7b3aecf7b9e96df08cc7da6e7be8df250c8a26a0ba8e9343617e1155abaa3dd01c4305a1e174f99f825a73b921c01c2a10ae7a06fc86722cfb755ffad531feef7990b0a42548e4205e5e2b63ffa9c608a95130d0080281dcba109d2c3e027fbcd1c2ce9c0"

temp = ""
result = []
t = 0
for i in range(len(a)):
    try:
        if b[i] == a[i]:
            temp += a[i]
            t = 0
        else:
            if t == 0:
                result.append(temp)
                temp = ""
                t = 1
    except:
        break

print(result)