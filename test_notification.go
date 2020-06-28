package pagerduty

// UpdateUser updates an existing user.
func (c *Client) SendTestNotification(headers map[string]string, userID string, contactMethodID string) error {
	resp, err := c.put("/users/"+userID+"/contact_methods/"+contactMethodID+"/send_test_notification", struct{}{}, &headers)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
