package g600

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

func TestNew(t *testing.T) {
	t.Run("Empty api token", func(t *testing.T) {
		c, err := New(Options{
			ApiKey: "",
		})

		require.Error(t, err, ErrEmptyApiKey)
		require.Nil(t, c)
	})
}

// Responses:
//
// Basic fee response:
// {"status":200,"recommendedMinimumFeeRatio":"12.81 Satoshi per byte","recommendedFeeInSatoshi":44015,"time":"Tue, 14 Aug 2018 19:08:18 GMT"}
func TestClient_RecommendedFee(t *testing.T) {
	c, err := New(Options{
		ApiKey: "token",
	})
	require.NoError(t, err)

	t.Run("Good response", func(t *testing.T) {
		defer gock.Off()

		gock.New(DefaultURL).
			Get("g600/api/v1/token/current-recommended-fee-ratio").
			Reply(200).
			BodyString(`{"status":200,"recommendedMinimumFeeRatio":"12.81 Satoshi per byte","recommendedFeeInSatoshi":44015,"time":"Tue, 14 Aug 2018 19:08:18 GMT"}`)

		fee, rec, err := c.RecommendedFee()
		require.NoError(t, err)
		require.Equal(t, int64(44015), fee)
		require.Equal(t, "12.81 Satoshi per byte", rec)
	})

	t.Run("Bad response", func(t *testing.T) {
		defer gock.Off()

		gock.New(DefaultURL).
			Get("g600/api/v1/token/current-recommended-fee-ratio").
			Reply(200).
			BodyString(`{"status":400,"recommendedMinimumFeeRatio":"12.81 Satoshi per byte","recommendedFeeInSatoshi":44015,"time":"Tue, 14 Aug 2018 19:08:18 GMT","message":"Hola"}`)

		fee, rec, err := c.RecommendedFee()
		require.Error(t, err)
		require.Equal(t, int64(0), fee)
		require.Equal(t, "", rec)
	})
}

// Responses:
//
// Basic transaction response:
// {"status":200,"message":{"Hash":"6784ffa910981359dee8a560925288649baab1a8a9233761bc3876fb67d4ae95","outputAddress":"3QYSQnXfdTuSGMFYRHMuKaevJC8rxsYmDG","status":"confirmed","scoreTime":"Tue Aug 14 2018 14:04:25 GMT+0000 (UTC)","agentId":"Primary","size":3436,"txValueUSD":112.36,"txValueBTC":0.01848737,"txValue$":112.36,"username":"nikolai@cryptopay.me"}}
//
// Unknown transaction response:
// {"status":400,"type":"TransactionNotFoundError","message":"Transaction not found!","Hash":"92db07c2a31b2677dffdf82467693c33eeaba5ced81edd6d9126c697703ab26b"}
//
// Bad API key:
// {"status":400,"type":"InvalidApiKeyError","message":"Invalid api key.","Hash":"92db07c2a31b2677dffdf82467693c33eeaba5ced81edd6d9126c697703ab26b"}
//
func TestClient_TransactionConfirm(t *testing.T) {
	c, err := New(Options{
		ApiKey: "token",
	})
	require.NoError(t, err)

	t.Run("Good response", func(t *testing.T) {
		defer gock.Off()

		gock.New(DefaultURL).
			Get("g600/api/v1/token/Primary/hash/output-address").
			Reply(200).
			BodyString(`{"status":200,"message":{"Hash":"6784ffa910981359dee8a560925288649baab1a8a9233761bc3876fb67d4ae95","outputAddress":"3QYSQnXfdTuSGMFYRHMuKaevJC8rxsYmDG","status":"confirmed","scoreTime":"Tue Aug 14 2018 14:04:25 GMT+0000 (UTC)","agentId":"Primary","size":3436,"txValueUSD":112.36,"txValueBTC":0.01848737,"txValue$":112.36,"username":"nikolai@cryptopay.me"}}`)

		res, err := c.TransactionConfirm("hash", "output-address")
		require.NoError(t, err)
		require.Equal(t, &TransactionConfirmationResponse{
			Hash:          "6784ffa910981359dee8a560925288649baab1a8a9233761bc3876fb67d4ae95",
			OutputAddress: "3QYSQnXfdTuSGMFYRHMuKaevJC8rxsYmDG",
			Username:      "nikolai@cryptopay.me",
			Status:        "confirmed",
			ScoreTime:     "Tue Aug 14 2018 14:04:25 GMT+0000 (UTC)",
			AgentID:       "Primary",
			Size:          3436,
			TxValueBTC:    0.01848737,
			TxValueUSD:    112.36,
			TxValue:       112.36,
		}, res)
	})

	t.Run("Bad response", func(t *testing.T) {
		defer gock.Off()

		gock.New(DefaultURL).
			Get("g600/api/v1/token/Primary/bad-hash/output-address").
			Reply(200).
			BodyString(`{"status":400,"type":"TransactionNotFoundError","message":"Transaction not found!","Hash":"92db07c2a31b2677dffdf82467693c33eeaba5ced81edd6d9126c697703ab26b"}`)

		res, err := c.TransactionConfirm("bad-hash", "output-address")
		require.Error(t, err)
		require.Nil(t, res)
	})
}

func TestClient_TransactionConfirmTestnet(t *testing.T) {
	c, err := New(Options{
		ApiKey: "token",
	})
	require.NoError(t, err)

	t.Run("Good response", func(t *testing.T) {
		defer gock.Off()

		gock.New(DefaultURL).
			Get("g600/apitest/v1/token/Primary/hash/output-address").
			Reply(200).
			BodyString(`{"status":200,"message":{"Hash":"6784ffa910981359dee8a560925288649baab1a8a9233761bc3876fb67d4ae95","outputAddress":"3QYSQnXfdTuSGMFYRHMuKaevJC8rxsYmDG","status":"confirmed","scoreTime":"Tue Aug 14 2018 14:04:25 GMT+0000 (UTC)","agentId":"Primary","size":3436,"txValueUSD":112.36,"txValueBTC":0.01848737,"txValue$":112.36,"username":"nikolai@cryptopay.me"}}`)

		res, err := c.TransactionConfirmTestnet("hash", "output-address")
		require.NoError(t, err)
		require.Equal(t, &TransactionConfirmationResponse{
			Hash:          "6784ffa910981359dee8a560925288649baab1a8a9233761bc3876fb67d4ae95",
			OutputAddress: "3QYSQnXfdTuSGMFYRHMuKaevJC8rxsYmDG",
			Username:      "nikolai@cryptopay.me",
			Status:        "confirmed",
			ScoreTime:     "Tue Aug 14 2018 14:04:25 GMT+0000 (UTC)",
			AgentID:       "Primary",
			Size:          3436,
			TxValueBTC:    0.01848737,
			TxValueUSD:    112.36,
			TxValue:       112.36,
		}, res)
	})

	t.Run("Bad response", func(t *testing.T) {
		defer gock.Off()

		gock.New(DefaultURL).
			Get("g600/apitest/v1/token/Primary/bad-hash/output-address").
			Reply(200).
			BodyString(`{"status":400,"type":"TransactionNotFoundError","message":"Transaction not found!","Hash":"92db07c2a31b2677dffdf82467693c33eeaba5ced81edd6d9126c697703ab26b"}`)

		res, err := c.TransactionConfirmTestnet("bad-hash", "output-address")
		require.Error(t, err)
		require.Nil(t, res)
	})
}

func TestClient_request(t *testing.T) {
	c, err := New(Options{
		ApiKey: "token",
	})
	require.NoError(t, err)

	t.Run("Bad URL provided", func(t *testing.T) {
		t.Skip()
	})

	t.Run("Client.Do error", func(t *testing.T) {
		t.Skip()
	})

	t.Run("Error decoding", func(t *testing.T) {
		defer gock.Off()

		gock.New(DefaultURL).
			Get("").
			Reply(200).
			BodyString(`i am a bad json`)

		_, err := c.request("", nil)
		require.Error(t, err)
	})

	t.Run("Error decoding error", func(t *testing.T) {
		defer gock.Off()

		gock.New(DefaultURL).
			Get("").
			Reply(200).
			BodyString(`{"status":400,"type":"TransactionNotFoundError","message":100.500,"Hash":"92db07c2a31b2677dffdf82467693c33eeaba5ced81edd6d9126c697703ab26b"}`)

		_, err := c.request("", nil)
		require.Error(t, err)
	})

	t.Run("Error decoding success", func(t *testing.T) {
		defer gock.Off()

		gock.New(DefaultURL).
			Get("").
			Reply(200).
			BodyString(`{"status":200,"message":100.500,"Hash":"92db07c2a31b2677dffdf82467693c33eeaba5ced81edd6d9126c697703ab26b"}`)

		_, err := c.request("", nil)
		require.Error(t, err)
	})
}
