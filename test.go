package main

import (
	"fmt"
	"bytes"
	// "time"
)

func main() {
	b := []byte("234rtyfdss")
	fmt.Printf("b %v \n", b)
	c := []byte{114, 116, 121, 102, 115}
	fmt.Printf("c %v \n", c)
	pb := bytes.Repeat(c, 3)
	fmt.Printf("pb %v \n", pb)
}


plaintext [212, 112, 50, 155, 212, 217, 55, 127, 163, 116, 68, 255, 225, 18, 61, 177, 101, 254, 5, 240, 136, 22, 132, 104, 100, 10, 161, 118, 24, 147, 13, 70, 194, 240, 15, 34, 123, 174, 2, 8, 187, 72, 114, 232, 38, 189, 140, 156, 249, 23, 68, 109, 143, 238, 218, 204, 142, 122, 215, 139, 188, 144, 135, 126, 139, 178, 98, 221, 253, 27, 194, 234, 72, 47, 20, 216, 158, 179, 106, 5, 35, 34, 154, 165, 82, 59, 62, 165, 184, 163, 118, 58, 214, 95, 84, 29, 163, 255, 87, 152, 237, 90, 129, 216, 156, 148, 218, 29, 34, 206, 79, 171, 216, 125, 190, 121, 95, 180, 3, 122, 254, 201, 235, 164, 188, 146, 192, 49, 142, 176, 91, 0, 176, 171, 240, 20, 67, 129, 46, 165, 41, 56, 61, 239] // 144
origData [212 112 50 155 212 217 55 127 163 116 68 255 225 18 61 177 101 254 5 240 136 22 132 104 100 10 161 118 24 147 13 70 194 240 15 34 123 174 2 8 187 72 114 232 38 189 140 156 249 23 68 109 143 238 218 204 142 122 215 139 188 144 135 126 139 178 98 221 253 27 194 234 72 47 20 216 158 179 106 5 35 34 154 165 82 59 62 165 184 163 118 58 214 95 84 29 163 255 87 152 237 90 129 216 156 148 218 29 34 206 79 171 216 125 190 121 95 180 3 122 254 201 235 164 188 146 192 49 142 176 91 0 176 171 240 20 67 129 46 165 41 56 61 239 16 16 16 16 16 16 16 16 16 16 16 16 16 16 16 16] 