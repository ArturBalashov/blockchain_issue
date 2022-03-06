# issue

Design and implement “Word of Wisdom” tcp server.

   - TCP server should be protected from DDOS attacks with the Prof of Work, the challenge-response protocol should be used.
   - The choice of the POW algorithm should be explained.
   - After Prof Of Work verification, the server should send one of the quotes from “Word of wisdom” book or any other collection of the quotes.
   - Docker file should be provided both for the server and for the client that solves the POW challenge
# Run

    docker-compose up -d
    docker-compose run client

# Why Hashcash

    This is the most known proof-of-work algorithm, which is used in Bitcoin's mining.
