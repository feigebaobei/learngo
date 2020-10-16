package main

import (
	"fmt"
	"bytes"
	"code.bonc.com.cn/blockchain/BONChain/common"
	"code.bonc.com.cn/blockchain/BONChain/common/hexutil"
	// "code.bonc.com.cn/blockchain/BONChain/crypto/sm"
	// "code.bonc.com.cn/blockchain/BONChain/crypto/sm/sm2"
	"code.bonc.com.cn/blockchain/BONChain/crypto/sm/sm4"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	// "encoding/hex"
	// "golang.org/x/crypto/sha3"
)


func main() {
	C_SM4Encrypt()
	C_SM4Decrypt()
	// fmd5()
}

// func fmd5 () {
// 	// data := "we45ty32345"
// 	// m := md5.Sum(data)
// 	fmt.Printf("hashCode %v \n", m)
// }
var (
	NULL = ""
)

func C_SM4Encrypt() string {
	// key := "96e79218965eb72c92a549dd5a330112"
	key := "0x96e79218965eb72c92a549dd5a330112"
	// data := "d470329bd4d9377fa37444ffe1123db165fe05f088168468640aa17618930d46c2f00f227bae0208bb4872e826bd8c9cf917446d8feedacc8e7ad78bbc90877e8bb262ddfd1bc2ea482f14d89eb36a0523229aa5523b3ea5b8a3763ad65f541da3ff5798ed5a81d89c94da1d22ce4fabd87dbe795fb4037afec9eba4bc92c031"
	data := "0xd470329bd4d9377fa37444ffe1123db165fe05f088168468640aa17618930d46c2f00f227bae0208bb4872e826bd8c9cf917446d8feedacc8e7ad78bbc90877e8bb262ddfd1bc2ea482f14d89eb36a0523229aa5523b3ea5b8a3763ad65f541da3ff5798ed5a81d89c94da1d22ce4fabd87dbe795fb4037afec9eba4bc92c031"
	// data := "{'privatekey':'01837f014db7fc5acd914f53839bdb5dbf4cd80ecbbb7bf966ba9619f34b627a'}"

	iv := make([]byte, 16)
	_, err := rand.Read(iv)
	if err != nil {
		fmt.Printf("err %v \n", err)
		return NULL
	}
	fmt.Printf("iv %v \n", iv)
	// [171 193 64 227 150 135 134 160 40 85 40 11 245 145 152 241]

	keyBytes, err := hexutil.Decode(key)
	if err != nil {
		fmt.Printf("err %v \n", err)
		return NULL
	}
	fmt.Printf("keyBytes %v \n", keyBytes)
	// [150 231 146 24 150 94 183 44 146 165 73 221 90 51 1 18]

	databytes, err := hexutil.Decode(data)
	if err != nil {
		fmt.Printf("err %v \n", err)
		return NULL
	}
	fmt.Printf("databytes %v \n", databytes)
	// [212 112 50 155 212 217 55 127 163 116 68 255 225 18 61 177 101 254 5 240 136 22 132 104 100 10 161 118 24 147 13 70 194 240 15 34 123 174 2 8 187 72 114 232 38 189 140 156 249 23 68 109 143 238 218 204 142 122 215 139 188 144 135 126 139 178 98 221 253 27 194 234 72 47 20 216 158 179 106 5 35 34 154 165 82 59 62 165 184 163 118 58 214 95 84 29 163 255 87 152 237 90 129 216 156 148 218 29 34 206 79 171 216 125 190 121 95 180 3 122 254 201 235 164 188 146 192 49]

	block, err := sm4.NewCipher(keyBytes)
	if err != nil {
		fmt.Printf("err %v \n", err)
		return NULL
	}
	fmt.Printf("block %v \n", &block)

	hashCode := md5.Sum(databytes)
	fmt.Printf("hashCode %v \n", hashCode)
	// [142 176 91 0 176 171 240 20 67 129 46 165 41 56 61 239]
	blockSize := block.BlockSize()
	plaintext := append(databytes, hashCode[:]...)
	fmt.Printf("plaintext %v \n", plaintext)
	// [212 112 50 155 212 217 55 127 163 116 68 255 225 18 61 177 101 254 5 240 136 22 132 104 100 10 161 118 24 147 13 70 194 240 15 34 123 174 2 8 187 72 114 232 38 189 140 156 249 23 68 109 143 238 218 204 142 122 215 139 188 144 135 126 139 178 98 221 253 27 194 234 72 47 20 216 158 179 106 5 35 34 154 165 82 59 62 165 184 163 118 58 214 95 84 29 163 255 87 152 237 90 129 216 156 148 218 29 34 206 79 171 216 125 190 121 95 180 3 122 254 201 235 164 188 146 192 49 142 176 91 0 176 171 240 20 67 129 46 165 41 56 61 239]

	origData := pkcs5Padding(plaintext, blockSize)
	fmt.Printf("origData %v \n", origData)
	// [212 112 50 155 212 217 55 127 163 116 68 255 225 18 61 177 101 254 5 240 136 22 132 104 100 10 161 118 24 147 13 70 194 240 15 34 123 174 2 8 187 72 114 232 38 189 140 156 249 23 68 109 143 238 218 204 142 122 215 139 188 144 135 126 139 178 98 221 253 27 194 234 72 47 20 216 158 179 106 5 35 34 154 165 82 59 62 165 184 163 118 58 214 95 84 29 163 255 87 152 237 90 129 216 156 148 218 29 34 206 79 171 216 125 190 121 95 180 3 122 254 201 235 164 188 146 192 49 142 176 91 0 176 171 240 20 67 129 46 165 41 56 61 239 16 16 16 16 16 16 16 16 16 16 16 16 16 16 16 16]

	blockMode := cipher.NewCBCEncrypter(block, iv)
	cryted := make([]byte, len(origData)+blockSize)
	fmt.Printf("cryted %v \n", cryted)
	block.Encrypt(cryted, iv)
	fmt.Printf("cryted %v \n", cryted)
	blockMode.CryptBlocks(cryted[blockSize:], origData)
	fmt.Printf("cryted %v \n", cryted)
	// [162 53 212 1 177 75 158 192 180 110 80 199 66 189 135 31 239 18 209 74 119 107 76 66 208 91 153 192 225 103 33 208 178 81 101 206 4 193 196 151 55 157 250 25 205 109 93 163 28 58 213 164 27 58 113 160 217 99 180 32 181 48 246 38 171 117 194 155 240 23 136 219 141 154 168 156 5 144 77 74 96 6 194 84 28 181 174 219 233 191 125 255 136 9 92 25 183 247 153 89 19 74 211 61 79 93 44 43 169 184 160 34 112 101 153 69 117 84 13 151 234 98 146 47 169 118 50 165 131 72 195 4 130 143 68 145 237 37 92 51 185 90 45 115 227 155 77 188 234 34 132 81 224 177 75 170 95 97 148 69 114 192 119 36 164 216 144 187 91 61 173 98 195 60 245 79]
	ct := hexutil.Encode(cryted)
	fmt.Printf("ct %v \n", ct)
	// 0xa235d401b14b9ec0b46e50c742bd871fef12d14a776b4c42d05b99c0e16721d0b25165ce04c1c497379dfa19cd6d5da31c3ad5a41b3a71a0d963b420b530f626ab75c29bf01788db8d9aa89c05904d4a6006c2541cb5aedbe9bf7dff88095c19b7f79959134ad33d4f5d2c2ba9b8a0227065994575540d97ea62922fa97632a58348c304828f4491ed255c33b95a2d73e39b4dbcea228451e0b14baa5f61944572c07724a4d890bb5b3dad62c33cf54f

	return ct

}

// pkcs5填充
func pkcs5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func C_SM4Decrypt () string {
	key := "0x96e79218965eb72c92a549dd5a330112"
	data := "0xd470329bd4d9377fa37444ffe1123db165fe05f088168468640aa17618930d46c2f00f227bae0208bb4872e826bd8c9cf917446d8feedacc8e7ad78bbc90877e8bb262ddfd1bc2ea482f14d89eb36a0523229aa5523b3ea5b8a3763ad65f541da3ff5798ed5a81d89c94da1d22ce4fabd87dbe795fb4037afec9eba4bc92c031"

	cipherText := common.FromHex(data)
	block, err := sm4.NewCipher(common.FromHex(key))
	if err != nil {
		return NULL
	}
	blockSize := block.BlockSize()
	iv := make([]byte, blockSize)
	block.Decrypt(iv, cipherText)

	blockMode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(cipherText)-blockSize)

	blockMode.CryptBlocks(plaintext, cipherText[16:])
	origText := pkcs5UnPadding(plaintext)
	if len(origText) == 0 {
		return NULL
	}
	hashCode := md5.Sum(origText[:len(origText)-md5.Size])
	if bytes.Equal(hashCode[:], origText[len(origText)-md5.Size:]) == true {
		mt := hexutil.Encode(origText[:len(origText)-md5.Size])
		fmt.Printf("mt %v \n", mt)
		return mt
	}
	return NULL

}

func pkcs5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	if unpadding > length {
		return src[:0]
	}
	return src[:(length - unpadding)]
}


