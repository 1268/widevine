package main

import (
   "41.neocities.org/widevine"
   "bytes"
   "encoding/base64"
   "encoding/hex"
   "encoding/json"
   "fmt"
   "net/http"
   "strconv"
)

func (d *drm_today) New(private_key, client_id []byte) error {
   key_id, err := hex.DecodeString(raw_key_id)
   if err != nil {
      return err
   }
   var pssh widevine.PsshData
   pssh.KeyIds = [][]byte{key_id}
   var module widevine.Cdm
   err = module.New(private_key, client_id, pssh.Marshal())
   if err != nil {
      return err
   }
   data, err := module.RequestBody()
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://lic.staging.drmtoday.com/license-proxy-widevine/cenc",
      bytes.NewReader(data),
   )
   if err != nil {
      return err
   }
   data, err = json.Marshal(map[string]string{
      "merchant": "client_dev",
      "userId":   "purchase",
   })
   if err != nil {
      return err
   }
   req.Header.Set("dt-custom-data", base64.StdEncoding.EncodeToString(data))
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   data, err = base64.StdEncoding.DecodeString(
      resp.Header.Get("x-dt-client-info"),
   )
   if err != nil {
      return err
   }
   err = json.Unmarshal(data, &d.client_info)
   if err != nil {
      return err
   }
   code, err := strconv.Atoi(resp.Header.Get("x-dt-resp-code"))
   if err != nil {
      return err
   }
   d.resp_code = resp_code(code)
   return nil
}

type drm_today struct {
   client_info client_info
   resp_code resp_code
}

// content.players.castlabs.com/demos/drm-agent/manifest.mpd
const raw_key_id = "6f6b1b9884f83d0b866a1bd8aca390d2"

type transport struct{}

func (transport) RoundTrip(req *http.Request) (*http.Response, error) {
   fmt.Println(req.URL)
   return http.DefaultTransport.RoundTrip(req)
}

type client_info struct {
   DrmVersion          *string
   HdcpSupport         string
   Manufacturer        string
   Model               string
   SecLevel            int64
   VmpStatus           string
   VrConstraintSupport bool
}

func (c *client_info) String() string {
   var data []byte
   if c.DrmVersion != nil {
      data = append(data, "drmVersion = "...)
      data = append(data, *c.DrmVersion...)
   }
   if data != nil {
      data = append(data, '\n')
   }
   data = append(data, "hdcpSupport = "...)
   data = append(data, c.HdcpSupport...)
   data = append(data, "\nmanufacturer = "...)
   data = append(data, c.Manufacturer...)
   data = append(data, "\nmodel = "...)
   data = append(data, c.Model...)
   data = append(data, "\nsecLevel = "...)
   data = strconv.AppendInt(data, c.SecLevel, 10)
   data = append(data, "\nvmpStatus = "...)
   data = append(data, c.VmpStatus...)
   data = append(data, "\nvrConstraintSupport = "...)
   data = strconv.AppendBool(data, c.VrConstraintSupport)
   return string(data)
}

type resp_code int

func (r resp_code) String() string {
   return "x-dt-resp-code = " + codes[r]
}

var codes = map[resp_code]string{
   0: "Success",
   1000: "General Internal Error",
   2000: "General Request Error",
   3000: "General Request Authentication Error",
   10001: "Bad Request",
   3e4: "General DRM error",
   4e4: "General Widevine Modular error",
   40001: "Widevine Device Certificate Revocation (wv 127)",
   40002: "Widevine Device Certificate Revocation - Permanently (wv 175)",
   41e3: "General Widevine Classic error",
   42e3: "General Playready error",
   43e3: "General Fairplay error",
   44e3: "General OMA error",
   44001: "OMA Device registration failed",
   45e3: "General CDRM error",
   45001: "CDRM Device registration failed",
   6e4: "CSL",
   60001: "CSL - INVALID",
   60100: "CSL - Denied by Stream Limiting",
   7e4: "General Output Protection",
   70001: "All keys filtered by EOP settings",
   8e4: "TAKE DOWN",
   80001: "TAKE DOWN - Denied by Take Down",
   9e4: "General GBL error",
   90001: "License delivery prohibited in your region",
}
