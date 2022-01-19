package auth

type Auth interface {
	Login(string, string) (string, int)
}
