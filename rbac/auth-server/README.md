# RBAC

Auth server demo implements [OAuth 2.0](https://tools.ietf.org/html/rfc6749).

## How to Generate Self-signed RSA Key Pairs

```
# openssl genrsa -out private_key.pem 2048
Generating RSA private key, 2048 bit long modulus
...........................+++
.............+++
e is 65537 (0x10001)

# openssl rsa -in private_key.pem -pubout -out public_key.pem
writing RSA key

# ls
private_key.pem  public_key.pem
```

**Note**: Make sure to generate the private key with PEM type of `RSA PRIVATE KEY` instead of `PRIVATE KEY`, 
otherwise will unable to decode private key PEM.

## How to Get Token

**Example Request**

```
POST http://localhost:8080/token?service=product&scope=fruits:apple:create,delete,update  HTTP/1.1

Headers:
	Authorization: Basic cHJvZHVjdC1tYXN0ZXI6MTIzNDU2 (Basic Auth for username `product-master` and password `123456`)
```

**Example Request**

```json
{
  "expires_in": 900,
  "issued_at": "2017-04-09T03:10:19Z",
  "token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzIwNDgiLCJraWQiOiJMTE1DOlhZVFc6S1BDSDpXNFJSOlFXNVo6V1RKQjo2Q09LOkI1V0w6UTI3MjpMVUlYOlZGU0c6SzRaSiJ9.eyJpc3MiOiJBdXRoIFNlcnZlciBEZW1vIiwic3ViIjoicHJvZHVjdC1tYXN0ZXIiLCJhdWQiOiJwcm9kdWN0IiwiZXhwIjoxNDkxNzA4MzE5LCJuYmYiOjE0OTE3MDc0MTgsImlhdCI6MTQ5MTcwNzQxOSwianRpIjoidDVXQTlnaHoyN3VwZlh5cCIsImFjY2VzcyI6W3sidHlwZSI6ImZydWl0cyIsIm5hbWUiOiJhcHBsZSIsImFjdGlvbnMiOlsiY3JlYXRlIiwiZGVsZXRlIiwidXBkYXRlIl19XX0.YOIvKoO02RcyAQbODl6EIi1p70muvLkfy7D-u-PIFlDh6JqaMrfyhYqbJh2t4jLOzij8NrrViJaCyrtG3ggEmhN6XZdNANtCJCNwkoHQQnNIfba-sdGN46_QwNAUQUhFutdN3plxVPiftaxzrlf7ffuqxnUwQpBn6Lqb_g8Gn18RXwGUMm5FIC3fcz3aHCHRkNcL6TTIPuxz8YDMYtCnH_QXpfTQOt7Owruh8OAYS2AVs-J9tWm0i2DGkd92TwE-FmmO2MwSYjwGI1PvmeWlqKpjfq2yx4H-gkYysQee2Ymd1f-6xp4FrFpu0LTBnriWZpwa0tm1Ar-4fA3AidL3xA"
}
```