package widevine

import (
   "154.pages.dev/protobuf"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
)

var tests = map[string]struct {
   response string
   url      string
   pssh protobuf.Message
}{
   "amc": {
      url:      "amcplus.com/movies/blackberry--1065021",
      response: "CAISmAUKLgoAEiYKEAV+drBU8k6ZsjsCuVMVA1IQATIQOlZidDX8oP4f+vmjvE8T8CABKAASHQgBEAAYACAAKAA4AEIASABQAFgAYAFwAHgBgAEAGkYSEDF6cQVXjcG7PpmOmfHvF3waMJuG+iumGx9AY5EwapkblVyBz8K9qvlStb63OxW5CwOVjHXUR0B/Rnmen0f5d7BkjCABGmIKEP3Bn0gybk/goXwKTwv51vsSEFeG0IyMLljwr4iwlfXLY+4aICqqO2jODZ0vt/CtXrT5f8CJmVIp3eLpNUZYD0j3j608IAIoATIICAAQKhgAIAA6CAgAECoYACAAYgJTRBplChAnFO2mertONbtRGq4FevFqEhA/YHhmaKtUsxON9Lomg7e7GiAehkKm44KyEeu6Mt3tXAc493bXo7rRqUXok0apb178ZyACKAEyCAgAECoYACAAOggIABAqGAAgAGIFQVVESU8aYgoQoJwVjRc0RZOwMd2/mA+Q5xIQf7w1bctJiq9AGdDeLMlQKhogmyO8ULGXSnJDpTuvmFSv47FkSEsMxUwsEfnQU0Y0ysMgAigBMggIARADGAAgADoICAAQKhgAIABiAkhEGmIKEOZvmO/LA0Pdp2Tz5UtJ51ISEAGPhXNkIfFasfpdbPdF04MaIPsewm3UvkM1/U3PDOVIwlHlLB2y3MMJflPyBb84ej/nIAIoATIICAAQAxgAIAA6CAgAECoYACAAYgJIRBpkChAMtku/LwFEA64yvxnrE5FqEhBEkBLE4DoGVgR7KNGuh4YNGiAlqzGfj5GEoGIEBY+eZZK09fOR770xxnsdWYdaFPLGDiACKAEyCAgBEAMYACAAOggIABAqGAAgAGIEVUhEMSDYiYKwBjgAGiAu3ZTSQypcLc1ermOWkafggihHQvKmnwnsJmTS8x79ViKAAo6gRwZ4aAzR9GPHMmst27OXvYtsV2r622JHL6Vt4ZsCqQpyHIlm4weCc3PDyTaN6QE7BmNCiDydtJsdcI+AXxToVvKLrWsYtPLsvRLOzPRFQDwICh8P98MhcA4Goa1Jwfzh/8KP0UrS76Plpy0apUXNdtTp4Wn9mE1n9wJIb5zsxUpwjHRVgmzG7hDGjjSR+iyImAHbvGUZ9SfvtB44mHuMHSGBgLltaXNa1bz6ujF7NAZP4OlqQ384bdVHtaK/IfEV9tF5zNPFfHtwkBh+yu3A1s7zA2/00+bcBggrZed4r3Q63G6/LoUtmTu3RiZsTfJf3f4LQEgWIUoWQtmGkQc6CAoGMTguMS4yQAFYAA==",
      pssh: protobuf.Message{
         protobuf.Field{Number:2, Type:2, Value:protobuf.Bytes("\xfd\xc1\x9fH2nO\xe0\xa1|\nO\v\xf9\xd6\xfb")},
         // optional
         protobuf.Field{Number:4, Type:2, Value:protobuf.Bytes("*")},
      },
   },
   "hulu": {
      response: "CAIS/QEKcAoAEmhZMlZyYzE4eU1XSTRNbVJqTWkxbFltSXlMVFJrTldFdFlUbG1PQzAyTXpGbU1EUTNNalkyTlRCZllXWmtZVFkxT0RRdFpUZGtPQzAwTVRJd0xUZzBZemd0WldFeE9EaGlPV1l5WXpSaiABKAASHwgBEAAYACAAKAAwADgAQgBIAFAAWABgAHABeAGAAQAaYAoQIbgtwuuyTVqp+GMfBHJmUBIQRwsFpNUILk3SEEWorgPnLxogi4r+vsi+BBQzOA+oovTTTbyoOlLnc7TNgCRhzAoLOU0gAigBMggIABAqGAAgADoICAAQKhgAIABiACCgnIKwBjgAGiC1Pi70ZIHRKIPjMed9cdvUc4QUA123+BghL3g5MmNxTCKAAgrENc/eO+LlZ3iWkSBXNJj1o9H15vuUit0DbLer4M6bhQLIMfufjT3YTaT85eRihlBPradY2lvwzA1g9lDtgGdYTJK7BKc0M+Mw7+nBiAQGMOvBUwretE/KDDLLTPsGCssK0pfac9CLzXStCEvs4SMOx1I1ACpGlL1ZGHg0XZowxk8wd0HvsZBVl+JBPZD7aBCl7AD67OJdK7KDsHaIKvBax8JYfDz7gRs2k3Hu+QG9pZWzng/e8q0uQu+e+wIhnocSbsLrk57eUinMPd2YJsYda3XbRSoq1Gq6CU7+dirFZhHB2eyY+smqh3w2sjSpSmhY7JvWqdrB1MFByW6K1MA6CAoGMTguMS4yQAFYAA==",
      url:      "hulu.com/watch/023c49bf-6a99-4c67-851c-4c9e7609cc1d",
      pssh: protobuf.Message{
         protobuf.Field{Number:2, Type:2, Value:protobuf.Bytes("!\xb8-\xc2\xeb\xb2MZ\xa9\xf8c\x1f\x04rfP")},
         // optional
         protobuf.Field{Number:4, Type:2, Value:protobuf.Bytes("*")},
      },
   },
   "mubi": {
      url:      "mubi.com/en/us/films/the-blair-witch-project",
      response: "CAISjQIKLgoAEiYKEJ7l+da2YE9ovVv0V6tyC1QQATIQOlZidDX8oP4f+vmjvE8T8CABKAASJwgBEAEYACCAmp4BKIDGCjCA4KgBOABCAEgAUABYAGABcAB4AYABABpGEhADSujgYl54uPBYb6K621imGjAEFFQIKQvIqAwhAHRpQDZi8rCrJKftUCM+oV9T2ay9wPoANvqXs9z23tgRcPISC/cgARpiChDq1Vx9mI1MbJY7KSqDl8oKEhABCiYvRwb6ouubj5DbxisvGiDqTmrKHxoSjGFISY3T+B54CaeUnaZTv6fOKuQMhl5bZCACKAEyCAgAECoYACAAOggIABAqGAAgAGICSEQg9qCCsAY4ABog0gVLsCkFctX3RPE/sf8pV/JfMaFJkKb8b6dWA9E3ZfwigAIMZ8jz7IRI/MJsJgUNmuGksQ2lhToXEysOZxjX1jIQuCtSgDpRWVueG52P8bOGZPh06WkY/Ab21XhNQJJFbq5KME8aTNehCpGCOXzjFkt137wWWCgGs9N78TuLKs+gbfClQHE01ztjUUlwj//4M1wSCFEDh6IcEpxf1Ma5mHJQOprlacG7dgK6Si/T8wIewJFCBlm0Hb/9ZKgjh4XBEfa9BY1tvdE81Y5gXCnu8SKF/wVqu7ccLj7rXODA207wcVQJxdgEojZbtxxgR4YMAI0V7SgM8okHIeWSlhFI6Tu3Wok6PinDGQrEX0N0rD1Jr3vy+y3dSnrKHf65wEY5hxr0OggKBjE4LjEuMkABWAA=",
      pssh: protobuf.Message{
         protobuf.Field{Number:2, Type:2, Value:protobuf.Bytes("\xea\xd5\\}\x98\x8dLl\x96;)*\x83\x97\xca\n")},
         // optional
         protobuf.Field{Number:4, Type:2, Value:protobuf.Bytes("*")},
      },
   },
   "nbc": {
      response: "CAISpwEKGAoAEhCj93lNr7YjRreQBYUorPULIAEoABIfCAEQABgAIAAoADAAOABCAEgAUABYAGAAcAB4AYABABpiChBkbsZgOexDcqdIT96VQdmaEhCnYc6PkBoDLbv1OTiG8L4bGiDWepd2Oeh/1Ifmi2qC7WySmxv6+DNSZRXPeTGDJH7vnCACKAEyCAgAECoYACAAOggIABAqGAAgAGICSEQgh/2BsAY4ABogsXHqxWzxEypKKi4seUCYpPeYx1CSifgoDTcryDEpP8AigAJtmGB+HCgsm56XVFd+vh1EG7gyDsphKbT8ZZyJkqDOYSDgDfeWEEj/WNXtcKUYp4x6rHktu/fGBAp5qL972ZLLebRazqh35X2NJYeJB3IP2c00mDt3YjOYlrDAxIQmbXLyFKTDTEEGmTyihbWbINETY3IiZcrTWRwzyNWjbtaNmE7bBju+/9Wriweziak7/42WMYYVQI7LGhLBWAK9EJXafwpVCiJSCItJ2Bc1YUvp1qCNqW4emu2Lk+SvYUSL0tpCqdwCqoqcux1j4HHM8qZm3I1TNrfQCTCllgjrgyvOVBUcOsDrLqTbLFoDPL0bgxKSiOO7ewGXJxbPtxEpTcC1OggKBjE4LjEuMkABWAA=",
      url:      "nbc.com/saturday-night-live/video/february-3-ayo-edebiri/9000283433",
      pssh:    protobuf.Message{
         protobuf.Field{Number:2, Type:2, Value:protobuf.Bytes("dn\xc6`9\xecCr\xa7HOޕAٚ")},
         // optional
         protobuf.Field{Number:4, Type:2, Value:protobuf.Bytes("*")},
      },
   },
   "paramount": {
      url:      "paramountplus.com/shows/video/bqsJh_z7o4AR6ktui_9y8wIHqzEEqbhr",
      response: "CAIS6gEKGgoAEghWyDISrU4VeiABKAA4AEAASJWFgrAGEhwIARAAGAAgACgAMAA4AEIASABQAFgAYAFwAHgBGkYSEN/ElvX3FkXDyR4w43+G4GUaMOE93lxPS4KZYOzVPt/544b4VuoHqDSsoOjT2cBLiU/fFE7yOh7uRJTCJ0n5a58ZxiABGl4KED3g8zwbik/Olh7aqVDi5zISEJwUy61dEZmKI7VNwJD8+T0aIEM2ON88GycAyYxcwSohr09QliqE9+8jk5oUB7Dx7q1qIAIoATIICAAQKhgAIAA6CAgAECoYACAAIJWFgrAGOAAaIG59jg/qd9AYwhVIJCPP8+yeG/361WernKxoCq8pmYq8IoACXH6ATVLEOrnHn1k21qqPHs8Np9iiM748cffYnDJaon71c+ioRoD4xPyFrvCpy3OmkYTPEltidhQ4ruMsJw0ig5reA//6qiAYdiOiafsN6ajzeQ5dxF15BypYMMU34TvHjZLoAS+KyoHesKMoTQYX3S1Hkxc0iqxP6TzzdWrLnKpg43SWmDQ2zDAmrv0ORLM7ZPyO7o9M0KzIE3szpuslTKZq3affPmvjrvEEdIg6Nyc4inlUDbDEnuE1rjCF6ge03EwMOIzrsCJbXgTF6hrqFAXkAWty483aDESQCVdIVH/UCj6jORYF9WW6umNxHD6+RLEDPUdlMakA768+KVudGjoICgYxNy40LjBAAVgA",
      pssh:     protobuf.Message{
         protobuf.Field{Number:2, Type:2, Value:protobuf.Bytes("=\xe0\xf3<\x1b\x8aOΖ\x1eکP\xe2\xe72")},
         // optional
         protobuf.Field{Number:4, Type:2, Value:protobuf.Bytes("*")},
      },
   },
   "peacock": {
      url:      "peacocktv.com/watch/playback/vod/GMO_00000000224510_02_HDSDR",
      response: "CAISvwEKLAoAEiQ4NjUzNDJjYi01NjUxLTRjOGItYjA4OC0wYjE1NWI0ZDUxZDAgASgAEh8IARAAGAAgACgAMAA4AEIASABQAFgAYABwAHgBgAEAGmIKEAAW4jRz6+d9k9jRpy3GkNcSEGPegnUYmDuqpgiKv43xs5UaIHH1LpRxUgM5BN/gA8o18nXgdY2l5vkCMhUfoMq2NGEdIAIoATIICAEQKhgAIAA6CAgBECoYACAAYgJTRCCNooKwBjjj3JWbBhog1Hfr9luk+PO9YTsecyiV3JdkxN7xfOlvduubmLy57qcigAI7y6Z/lYpUTJAjpNQYdJkgWxcTbzoS1ifi27o1p8n3fLWDdTi2TrhBdhEEVca5M0ZtCUR492ZSYVPUtcnl7n4d56SVS3RwrFazI4xpsGBdUb8gpqqvmAV87XQY73EUdgG/c9Z+JQuECFEDDrFeiTvQLKwdfpqBekI3K0VA/AWq+fsnpvLkmvAQXX2TLBAIH266eGOz32Xu1r/NrKsg9CweD2/RWQpljSDwWA9l/RWslRQuTDlZdqy5kbihrOzLkQw9DJetx3rY59ENVMTrqQ+LE4Jld83bsN3onT3JulngcsH9Qh5rVtSXsShPBUqZLdhfh4HrEAV6Av5l22iz8Dx2OggKBjE4LjEuMkABWAA=",
      pssh:     protobuf.Message{
         protobuf.Field{Number:2, Type:2, Value:protobuf.Bytes("\x00\x16\xe24s\xeb\xe7}\x93\xd8ѧ-Ɛ\xd7")},
         // optional
         protobuf.Field{Number:4, Type:2, Value:protobuf.Bytes("*")},
      },
   },
   "roku": {
      response: "CAISpAEKGgoAEgjX7INTgEMsMyABKAA4AEAASLuDgrAGEh4IARAAGAAgACgAMIDGCjgAQgBIAFAAWABgAHAAeAEaXgoQvfpNbNs5cC5baB+QYX+afhIQI5DiUKryEwPFVPj6j1sAKBogyKXm2gclc0YjMDLPF0Jh0UruSgpEUK6qlR3+e7zk48IgAigBMggIABAqGAAgADoICAAQKhgAIAAgu4OCsAY4ABogvFIfvRg3GuuL6G5fKT9WSYEuMQUSfrq3xI2dnhfndPcigAJ9qrQ8SKM5LHGUqwTLzf5irFBVQdjy6GO8B4X9/uhdSGGQdhUstrf1GNuzJ5KB8I9najUvp0JVfmyHAIQ5HEtThPdZkp17huho+zUjzn9nKS6h4t3BT2oIAj2TI9pYCN+qGZupuHwnVj6hYTwZ0SCePVj4BDzmd54L3aE3LHhO9QAFczZkYD7cw2+JzRzGb1Z2T+u7nzKMeivY+Why7/WA5d++JcC0U/YUUqChGa41tA8o1Ky3muIeY1C4gxTE8VH73Ofb5yABTPtySnoyUftz84ooKsuF4fqKG8O+959lPDy/kVR42XxHHiTONYnpv4WfkHr9CsfvuKF12pvE8TBoOggKBjE3LjAuMUABWAA=",
      url:      "therokuchannel.roku.com/watch/105c41ea75775968b670fbb26978ed76",
      // pssh optional, its embedded in the request URL
      pssh:     protobuf.Message{
         protobuf.Field{Number:2, Type:2, Value:protobuf.Bytes("\xbd\xfaMl\xdb9p.[h\x1f\x90a\x7f\x9a~")},
         // 2024-3-31 any single character is fine
         protobuf.Field{Number:4, Type:2, Value:protobuf.Bytes("*")},
      },
   },
}

func TestResponse(t *testing.T) {
   for _, test := range tests {
      module, err := new_module(test.pssh)
      if err != nil {
         t.Fatal(err)
      }
      signed, err := base64.StdEncoding.DecodeString(test.response)
      if err != nil {
         t.Fatal(err)
      }
      license, err := module.response(signed)
      if err != nil {
         t.Fatal(err)
      }
      key, ok := module.Key(license)
      if !ok {
         t.Fatal("CDM.Key")
      }
      fmt.Println(test.url)
      fmt.Printf("%x\n\n", key)
   }
}

func new_module(pssh protobuf.Message) (*CDM, error) {
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
   var protect PSSH
   protect.Data = pssh.Encode()
   if err := protect.Consume(); err != nil {
      return nil, err
   }
   return protect.CDM(private_key, client_id)
}
