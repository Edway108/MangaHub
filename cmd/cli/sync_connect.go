package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func SyncConnect(token string) error {
    fmt.Println("Connecting to TCP sync server at localhost:9090...")

    conn, err := net.Dial("tcp", "localhost:9090")
    if err != nil {
        return err
    }

    encoder := json.NewEncoder(conn)
    decoder := json.NewDecoder(conn)

    // send auth
    encoder.Encode(map[string]string{
        "type":  "auth",
        "token": token,
    })

    var resp map[string]interface{}
    if err := decoder.Decode(&resp); err != nil {
        return err
    }

    if resp["type"] != "auth_ok" {
        return fmt.Errorf("auth failed")
    }

    fmt.Println(" Connected successfully!")
    fmt.Println("Session ID:", resp["session_id"])

    return nil
}
