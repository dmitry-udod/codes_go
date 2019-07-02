package tests

import (
	"encoding/json"
	"github.com/dmitry-udod/codes_go/cmd"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFopXmlStringDecode(t *testing.T) {

	recordStr := `<RECORD><FIO>САЄНКО ОЛЬГА СЕРГІЇВНА</FIO><ADDRESS>91007, Луганська обл., місто Луганськ, Артемівський район, ВУЛИЦЯ ПРИВОЗНА, будинок 55</ADDRESS><KVED>47.82 Роздрібна Торгівля з лотків і на ринках текстильними виробами, одягом і взуттям</KVED><STAN>припинено</STAN></RECORD>`

	record := cmd.DecodeFopXmlString(recordStr)

	assert.Equal(t, "Саєнко Ольга Сергіївна", record.FullName)
	assert.Equal(t, "91007, Луганська Обл., Місто Луганськ, Артемівський Район, Вулиця Привозна, Будинок 55", record.Address)
	assert.Equal(t, "47.82 роздрібна торгівля з лотків і на ринках текстильними виробами, одягом і взуттям", record.Activity)
	assert.Equal(t, "припинено", record.Status)
	assert.Equal(t, TEST_DOCUMENT_ID, record.GenerateId())

	data, err := json.Marshal(record)
	assert.Nil(t, err)
	assert.Equal(t, `{"full_name":"Саєнко Ольга Сергіївна","address":"91007, Луганська Обл., Місто Луганськ, Артемівський Район, Вулиця Привозна, Будинок 55","activity":"47.82 роздрібна торгівля з лотків і на ринках текстильними виробами, одягом і взуттям","status":"припинено"}`, string(data))
}
