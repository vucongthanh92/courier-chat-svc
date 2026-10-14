package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	utilsError "github.com/vucongthanh92/go-base-utils/helper/http_error"
)

func IsUniqueViolation(errBuilder *utilsError.ErrorBuilder) bool {
	if errBuilder == nil || errBuilder.LogError == nil {
		return false
	}

	message := strings.ToLower(errBuilder.LogError.Error())
	return strings.Contains(message, "sqlstate 23505") ||
		strings.Contains(message, "duplicate key")
}

// GetHeaderFromKey retrieves a specific header value from the context using the provided key.
func GenerateConversationDirectKey(memberUserIDs []uint64) string {
	if len(memberUserIDs) == 0 {
		return ""
	}

	ids := append([]uint64(nil), memberUserIDs...)
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })

	parts := make([]string, 0, len(ids))
	for _, id := range ids {
		parts = append(parts, strconv.FormatUint(id, 10))
	}

	canonical := strings.Join(parts, ":")
	sum := sha256.Sum256([]byte(canonical))
	return hex.EncodeToString(sum[:])
}

func GenerateSystemConversationDirectKey(userID uint64, name string) string {
	if userID == 0 || strings.TrimSpace(name) == "" {
		return ""
	}
	canonical := fmt.Sprintf("system:%d:%s", userID, strings.TrimSpace(name))
	sum := sha256.Sum256([]byte(canonical))
	return hex.EncodeToString(sum[:])
}

// func NormalizeMemberIDs normalizes a slice of member IDs by removing duplicates and ensuring all IDs are greater than 0.
func NormalizeMemberIDs(ids []uint64) ([]uint64, error) {

	if len(ids) == 0 {
		return nil, errors.New("member_user_ids is required")
	}

	seen := make(map[uint64]struct{}, len(ids))
	out := make([]uint64, 0, len(ids))
	for _, id := range ids {
		_, isExist := seen[id]
		switch {
		case id <= 0:
			return nil, errors.New("member user id must be greater than 0")
		case isExist:
			continue
		default:
			seen[id] = struct{}{}
			out = append(out, id)
		}
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i] < out[j]
	})

	return out, nil
}
