package wbyandex

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/billing/v1"
)

func dataSourceWebbankirYandexBillingAccountContent() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceWebbankirYandexBillingAccountContentRead,
		Schema: map[string]*schema.Schema{
			"billing_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"balance": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceWebbankirYandexBillingAccountContentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	ctx := config.Context()

	billingId := d.Get("billing_id").(string)
	d.SetId(billingId)

	data, err := config.sdk.Billing().BillingAccount().Get(ctx, &billing.GetBillingAccountRequest{Id: billingId})
	//data, err := config.sdk.CertificatesData().CertificateContent().Get(ctx, &certificatemanager.GetCertificateContentRequest{CertificateId: certificateId })
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("Billing Id %q", d.Id()))
	}

	if err := d.Set("balance", data.Balance); err != nil {
		return err
	}
	//if err := d.Set("certificate_chain", strings.Join(data.CertificateChain, "\n")); err != nil {
	//	return err
	//}

	return d.Set("billing_id", billingId)
}
