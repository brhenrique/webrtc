package ice

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	hostCandidatePreference  uint16 = 126
	srflxCandidatePreference uint16 = 100
)

// Candidate represents an ICE candidate
type Candidate interface {
	GetBase() *CandidateBase
	String(component int) string
}

// CandidateBase represents an ICE candidate, a base with enough attributes
// for host candidates, see CandidateSrflx and CandidateRelay for more
type CandidateBase struct {
	Protocol TransportType
	Address  string
	Port     int
	LastSeen time.Time
}

func (c *CandidateBase) priority(typePreference uint16, component uint16) uint16 {
	localPreference := uint16(rand.Uint32() / 2)
	return (2^24)*typePreference +
		(2^8)*localPreference +
		(2^0)*(256-component)
}

// CandidateHost is a Candidate of typ Host
type CandidateHost struct {
	CandidateBase
}

// String for CandidateHost
func (c *CandidateHost) String(component int) string {
	return fmt.Sprintf("udpcandidate %d udp %d %s %d typ host generation 0",
		component, c.CandidateBase.priority(hostCandidatePreference, uint16(component)), c.CandidateBase.Address, c.CandidateBase.Port)
}

// GetBase returns the CandidateBase, attributes shared between all Candidates
func (c *CandidateHost) GetBase() *CandidateBase {
	return &c.CandidateBase
}

// Address for CandidateHost
func (c *CandidateHost) Address() string {
	return c.CandidateBase.Address
}

// Port for CandidateHost
func (c *CandidateHost) Port() int {
	return c.CandidateBase.Port
}

// CandidateSrflx is a Candidate of typ Server-Reflexive
type CandidateSrflx struct {
	CandidateBase
	RemoteAddress string
	RemotePort    int
}

// String for CandidateSrflx
func (c *CandidateSrflx) String(component int) string {
	return fmt.Sprintf("udpcandidate %d udp %d %s %d typ srflx raddr %s rport %d generation 0",
		component, c.CandidateBase.priority(srflxCandidatePreference, uint16(component)), c.CandidateBase.Address, c.CandidateBase.Port, c.RemoteAddress, c.RemotePort)
}

// GetBase returns the CandidateBase, attributes shared between all Candidates
func (c *CandidateSrflx) GetBase() *CandidateBase {
	return &c.CandidateBase
}