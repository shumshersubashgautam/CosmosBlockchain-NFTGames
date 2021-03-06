package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/singhp1069/spunky/x/spunky/types"
)

var _ = strconv.Itoa(0)

func CmdTransferNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-nft [nft] [recipient]",
		Short: "Broadcast message transferNFT",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			nft, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argRecipient := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTransferNFT(
				clientCtx.GetFromAddress().String(),
				argRecipient,
				nft,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
