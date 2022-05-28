package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultGenesisIsCorrect(t *testing.T) {
	require.EqualValues(t,
		&GenesisState{
			StoredGameList: []StoredGame{},
			NextGame:       &NextGame{uint64(1)},
		},
		DefaultGenesis())
}

/*
func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				NextGame: &types.NextGame{
					IdValue: 23,
				},
				NextGame: &types.NextGame{
					IdValue: 68,
				},
				NextGame: &types.NextGame{
					IdValue: 37,
				},
				NextGame: &types.NextGame{
					IdValue: 31,
				},
				StoredGameList: []types.StoredGame{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated storedGame",
			genState: &types.GenesisState{
				StoredGameList: []types.StoredGame{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}*/
