package widevine

import (
   "154.pages.dev/encoding/protobuf"
   "bytes"
   "crypto/rsa"
   "crypto/x509"
   "encoding/hex"
   "encoding/pem"
   "io"
   "net/http"
)

// some videos require key_id and content_id, so entire PSSH is needed
func New_Module(private_key, client_ID, pssh []byte) (*Module, error) {
   block, _ := pem.Decode(private_key)
   var (
      err error
      mod Module
   )
   mod.private_key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
   if err != nil {
      return nil, err
   }
   var m protobuf.Message
   m.Add_Bytes(1, client_ID)
   m.Add(2, func(m *protobuf.Message) { // ContentId
      m.Add(1, func(m *protobuf.Message) { // CencId
         m.Add_Bytes(1, pssh[32:])
      })
   })
   mod.license_request = m.Append(nil)
   return &mod, nil
}

func unpad(buf []byte) []byte {
   if len(buf) >= 1 {
      pad := buf[len(buf)-1]
      if len(buf) >= int(pad) {
         buf = buf[:len(buf)-int(pad)]
      }
   }
   return buf
}

func (c Container) String() string {
   return hex.EncodeToString(c.Key)
}

type Containers []Container

func (c Containers) Content() *Container {
   for _, container := range c {
      if container.Type == 2 {
         return &container
      }
   }
   return nil
}

type Module struct {
   license_request []byte
   private_key *rsa.PrivateKey
}
func (m Module) Post(post Poster) (Containers, error) {
   signed_request, err := m.signed_request()
   if err != nil {
      return nil, err
   }
   body, err := post.Request_Body(signed_request)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", post.Request_URL(), bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   if head := post.Request_Header(); head != nil {
      req.Header = head
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   body, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   body, err = post.Response_Body(body)
   if err != nil {
      return nil, err
   }
   return m.signed_response(body)
}

type Poster interface {
   Request_URL() string
   Request_Header() http.Header
   Request_Body([]byte) ([]byte, error)
   Response_Body([]byte) ([]byte, error)
}

type no_operation struct{}

func (no_operation) Read(buf []byte) (int, error) {
   return len(buf), nil
}

