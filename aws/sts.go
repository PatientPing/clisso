package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func AssumeSAMLRole(PrincipalArn, RoleArn, SAMLAssertion string) (*Credentials, error) {
	input := sts.AssumeRoleWithSAMLInput{
		PrincipalArn:  aws.String(PrincipalArn),
		RoleArn:       aws.String(RoleArn),
		SAMLAssertion: aws.String(SAMLAssertion),
	}

	sess := session.Must(session.NewSession())
	svc := sts.New(sess)

	aResp, err := svc.AssumeRoleWithSAML(&input)
	if err != nil {
		return nil, fmt.Errorf("assuming role: %v", err)
	}

	keyID := *aResp.Credentials.AccessKeyId
	secretKey := *aResp.Credentials.SecretAccessKey
	sessionToken := *aResp.Credentials.SessionToken
	expiration := *aResp.Credentials.Expiration

	creds := Credentials{
		AccessKeyID:     keyID,
		SecretAccessKey: secretKey,
		SessionToken:    sessionToken,
		Expiration:      expiration,
	}

	return &creds, nil
}
