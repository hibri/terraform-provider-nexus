package nexus


type IQConnection struct {
	Enabled             bool   `json:"enabled",bool`
	ShowLink            bool   `json:"showLink",bool`
	Url                 string `json:"url",string`
	AuthenticationType  string `json:"authenticationType",string`
	Username            string `json:"username",string`
	Password            string `json:"password",string`
	UseTrustStoreForUrl bool   `json:"useTrustStoreForUrl",bool`
	TimeoutSeconds      int    `json:"timeoutSeconds",int`
	Properties          string `json:"properties",string`

}

type IQConnectionClient struct {
	client              NexusClient
	endPoint string
}

func NewIQConnectionClient(client NexusClient) IQConnectionClient{
  return IQConnectionClient{ client : client, endPoint:  "/beta/iq"}
}
func (iqc IQConnectionClient ) Get() (IQConnection, error) {
	iq := IQConnection{}
	err := iqc.client.Get(&iq, iqc.endPoint)
	return iq, err
}

func (iqc IQConnectionClient) Update(iq IQConnection) (IQConnection, error) {

	_ , err := iqc.client.Update(iq, iqc.endPoint)
	if err == nil {
		return iqc.Get()
	}
	return iq, err
}
