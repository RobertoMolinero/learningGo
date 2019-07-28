# SSL Encryption

## Configuration

### Step 1: Configure the hostname of the server

```
SERVER_CN=localhost
```

### Step 2: Generate Certificate Authority (*.key) + Trust Certificate (*.crt)

```
openssl genrsa -passout pass:secret -des3 -out ca.key 4096
openssl req -passin pass:secret -new -x509 -days 365 -key ca.key -out ca.crt -subj "/CN=${SERVER_CN}"
```

### Step 3: Generate the Server Private Key (server.key) 

```
openssl genrsa -passout pass:secret -des3 -out server.key 4096
```

### Step 4: Get a certificate signing request from the CA (server.csr)

```
openssl req -passin pass:secret -new -key server.key -out server.csr -subj "/CN=${SERVER_CN}"
```

### Step 5: Sign the certificate with the CA we created (self signing) - server.crt

```
openssl x509 -req -passin pass:secret -days 365 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt
```

### Step 6: Convert the server certificate to .pem format (server.pem) - usable by gRPC

```
openssl pkcs8 -topk8 -nocrypt -passin pass:secret -in server.key -out server.pem
```