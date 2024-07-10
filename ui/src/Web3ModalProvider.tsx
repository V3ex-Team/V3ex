import { createWeb3Modal } from '@web3modal/wagmi/react';
import { defaultWagmiConfig } from '@web3modal/wagmi/react/config';
import { WagmiProvider } from 'wagmi';
import { linea, lineaSepolia } from 'wagmi/chains';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

// 0. Setup queryClient
const queryClient = new QueryClient();

// 1. Get projectId from https://cloud.walletconnect.com
const projectId = 'fd50cb151f14f8ba05db1f56bfd4aa62';

// 2. Create wagmiConfig
const metadata = {
  name: 'Web3Modal',
  description: 'Web3Modal Example',
  url: 'https://web3modal.com', // origin must match your domain & subdomain
  icons: ['https://avatars.githubusercontent.com/u/37784886'],
};

const chains = [lineaSepolia, linea] as const;
const config = defaultWagmiConfig({
  chains,
  projectId,
  metadata,
  enableCoinbase: false,
});

// 3. Create modal
createWeb3Modal({
  wagmiConfig: config,
  projectId,

  defaultChain: linea,
  enableAnalytics: true, // Optional - defaults to your Cloud configuration
  enableOnramp: true, // Optional - false as default

  allWallets: 'HIDE',
});

export function Web3ModalProvider({ children }) {
  return (
    <WagmiProvider config={config}>
      <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
    </WagmiProvider>
  );
}

export default function ConnectButton() {
  return (
    <div style={{ marginRight: '16px' }}>
      <w3m-button balance="hide" size="sm" />
    </div>
  );
}
