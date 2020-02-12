package dto

//MerklePathDto MerklePathDto
type MerklePathDto struct {
	MerklePathNodes []MerklePathNodeDto
}

//MerklePathNodeDto MerklePathNodeDto
type MerklePathNodeDto struct {
	Hash            string
	IsLeftChildNode bool
}