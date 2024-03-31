package widevine

import (
   "bufio"
   "bytes"
   "encoding/json"
   "errors"
   "fmt"
   "io"
   "net/http"
   "os"
   "testing"
)

func TestPeacock(t *testing.T) {
   key, err := request("peacock", nil)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

func TestHulu(t *testing.T) {
   key, err := request("hulu", nil)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

func TestAmc(t *testing.T) {
   key, err := request("amc", nil)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

func TestParamount(t *testing.T) {
   key, err := request("paramount", nil)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

func TestMubi(t *testing.T) {
   unwrap := func(b []byte) ([]byte, error) {
      var s struct {
         License []byte
      }
      err := json.Unmarshal(b, &s)
      if err != nil {
         return nil, err
      }
      return s.License, nil
   }
   key, err := request("mubi", unwrap)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

func TestRoku(t *testing.T) {
   key, err := request("roku", nil)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

func TestNbc(t *testing.T) {
   key, err := request("nbc", nil)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

type unwrapper func([]byte) ([]byte, error)

func request(name string, unwrap unwrapper) ([]byte, error) {
   home, err := os.UserHomeDir()
   if err != nil {
      return nil, err
   }
   private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
   if err != nil {
      return nil, err
   }
   client_id, err := os.ReadFile(home + "/widevine/client_id.bin")
   if err != nil {
      return nil, err
   }
   file, err := os.Open("testdata/" + name + ".bin")
   if err != nil {
      return nil, err
   }
   defer file.Close()
   req, err := http.ReadRequest(bufio.NewReader(file))
   if err != nil {
      return nil, err
   }
   var protect PSSH
   protect.Data = tests[name].pssh.Encode()
   protect.m = tests[name].pssh
   module, err := protect.CDM(private_key, client_id)
   if err != nil {
      return nil, err
   }
   body, err := module.request_signed()
   if err != nil {
      return nil, err
   }
   req.Body = io.NopCloser(bytes.NewReader(body))
   req.ContentLength = 0
   req.Header.Del("accept-encoding")
   req.RequestURI = ""
   req.URL.Host = req.Host
   req.URL.Scheme = "https"
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b bytes.Buffer
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   body, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   if unwrap != nil {
      body, err = unwrap(body)
      if err != nil {
         return nil, err
      }
   }
   license, err := module.response(body)
   if err != nil {
      return nil, err
   }
   key, ok := module.Key(license)
   if !ok {
      return nil, errors.New("CDM.Key")
   }
   res.Write(os.Stdout)
   return key, nil
}
