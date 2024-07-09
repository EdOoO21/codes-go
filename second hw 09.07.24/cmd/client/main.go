package main

import (
	"awesomeProject/accounts/dto"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

type Command struct {
	Port   int
	Host   string
	Cmd    string
	Name   string
	Amount int
}

func main() {
	portVal := flag.Int("port", 1323, "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	amountVal := flag.Int("amount", 0, "amount of account")

	flag.Parse()

	cmd := Command{
		Port:   *portVal,
		Host:   *hostVal,
		Cmd:    *cmdVal,
		Name:   *nameVal,
		Amount: *amountVal,
	}

	if err := do(cmd); err != nil {
		panic(err)
	}
}

func do(cmd Command) error {
	switch cmd.Cmd {
	case "create":
		if err := create(cmd); err != nil {
			return fmt.Errorf("create account failed: %w", err)
		}

		return nil
	case "get":
		if err := get(cmd); err != nil {
			return fmt.Errorf("get account failed: %w", err)
		}

		return nil
	case "patch":
		if err := patch(cmd); err != nil {
			return fmt.Errorf("patch account failed: %w", err)
		}
		return nil
	case "delete":
		if err := delete(cmd); err != nil {
			return fmt.Errorf("delete account failed: %w", err)
		}
		return nil
	case "change":
		if err := change(cmd); err != nil {
			return fmt.Errorf("change account failed: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("unknown command %s", cmd.Cmd)
	}
}

func create(cmd Command) error {
	request := dto.CreateAccountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/create", cmd.Host, cmd.Port),
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusCreated {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func get(cmd Command) error {
	resp, err := http.Get(
		fmt.Sprintf("http://%s:%d/account?name=%s", cmd.Host, cmd.Port, cmd.Name),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}

		return fmt.Errorf("resp error %s", string(body))
	}

	var response dto.GetAccountResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("json decode failed: %w", err)
	}

	fmt.Printf("response account name: %s and amount: %d", response.Name, response.Amount)

	return nil
}

func delete(cmd Command) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://%s:%d/account/delete?name=%s", cmd.Host, cmd.Port, cmd.Name), nil)
	if err != nil {
		return fmt.Errorf("http delete failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("http delete failed: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode == http.StatusCreated {
		return nil
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}
	return fmt.Errorf("resp error %s", string(body))
}

func change(cmd Command) error {
	request := dto.CreateAccountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	}
	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("http://%s:%d/account/change", cmd.Host, cmd.Port), bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("http patch failed: %w", err)

	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("http change name failed: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode == http.StatusCreated {
		return nil
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}
	return fmt.Errorf("resp error %s", string(body))
}

func patch(cmd Command) error {
	request := dto.CreateAccountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	}
	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}
	req, err := http.NewRequest("PATCH",
		fmt.Sprintf("http://%s:%d/account/patch", cmd.Host, cmd.Port),
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http patch amount failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("http patch amount failed: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode == http.StatusCreated {
		return nil
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}
	return fmt.Errorf("resp error %s", string(body))

}
