rm *.pem
# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=KZ/ST=Almaty region/L=Almaty/O=23m/OU=Commercial org/CN=*.23m.kz/emailAddress=23m@gmail.com"

echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text

# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=KZ/ST=Mangistau region/L=Aktau/O=pc-book/OU=Computer/CN=*.pcbook.kz/emailAddress=pcbook@gmail.com"
# 3. Use CA's private key to sign web server's CSR and get back signed certificates
openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf

echo "Server's signed certificate"
openssl x509 -in server-cert.pem -noout -text
# 4. Verify the certificate:
# openssl verify -CAfile ca-cert.pem server-cert.pem

# 5. Generate web client's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout client-key.pem -out client-req.pem -subj "/C=KZ/ST=Akmola region/L=Nursultan/O=pc-client/OU=Computer/CN=*.pcclient.kz/emailAddress=pcclient@gmail.com"
# 6. Use CA's private key to sign web client's CSR and get back signed certificates
openssl x509 -req -in client-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem -extfile client-ext.cnf

echo "Client's signed certificate"
openssl x509 -in client-cert.pem -noout -text