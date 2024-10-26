// idl/git-typo.thrift
namespace go typo

struct GitTypoReq {
    1: string GitLink  (go.tag =  'json:"gitLink" query:"gitLink"');
    2: string Config (api.query="config");
}

// struct GitTypoReqResp {
//     1: string RespBody;
// }


service GitTypoService {
    string GitTypoMethod(1: GitTypoReq request) (api.post="/git-typo");
}

