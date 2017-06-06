# Yo Consul! A Hashicorp Consul Catalog client

This is a command line client which provides basic admin interaction with a local Hashicorp Consul Catalog - I wrote it because I was getting annoyed having to maintain JSON files and uploading them manually to Consul for any services that couldn't register themselves.

## Usage

* Registering a service: 
`yoc register --id service_id --name service_name --address address_of_service --port 8080 --tag "Tags for use by service"`

* Deregistering a service:
`yoc deregister service_id`

* Listing registered services:
`yoc services`

## Limitations

In this initial release you can only talk to the local (127.0.0.1 or localhost) Consul server on the default port and there is no authentication involved.  These things are coming.
