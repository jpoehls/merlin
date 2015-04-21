package main

var helpCert = &Command{
	UsageLines: []string{"cert"},
	Short:      "getting a conduit certificate",
	Long: `
conduit.connect is used to establish an authenticated session
with Phabricator. This requires a valid username and certificate.

To get your conduit certificate, go to:
  
  https://{MY_PHABRICATOR_HOST}/settings/panel/conduit

and copy/paste the certificate you find there.
	`,
}
