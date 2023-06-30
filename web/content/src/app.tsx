import { useEffect, useState } from 'react'
import "@massalabs/react-ui-kit/src/global.css";
import { Button, Dropdown } from "@massalabs/react-ui-kit"

interface WalletsProps {
  wallets: null | Array<string>,
  setWallets: (wallets: Array<string>) => void
}

const Wallets = ({ setWallets, wallets }: WalletsProps) => {
  const getWallets = async () => {
    var myHeaders = new Headers();
    myHeaders.append("Accept", "application/json");

    var requestOptions = {
      method: 'GET',
      headers: myHeaders,
    };

    fetch("https://station.massa/plugin/massa-labs/massa-wallet/api/accounts", requestOptions)
      .then(response => response.json())
      .then(result => setWallets(result.map((wallet: any) => wallet.nickname)))
      .catch(error => { console.log(error); setWallets([]) });
  }
  useEffect(() => {
    if (wallets === null) {
      getWallets();
    }
  });
  if (wallets === null) {
    return (
      <>
        <div>
          <h1 className="text-3xl font-bold">
            Loading...
          </h1>
        </div>
      </>
    )
  } else if (wallets.length === 0) {
    return (
      <>
        <div>
          <h1 className="text-neutral">
            No wallet found. Please install the wallet plugin provided by Massa Labs and create a wallet to use this plugin.
          </h1>
        </div>
      </>
    )
  } else {
    // Dropdown
    return (
      <>
        <Dropdown options={wallets.map(wallet => ({ item: wallet }))} />
      </>
    )
  }
}


export function App() {
  const [wallets, setWallets] = useState<null | Array<string>>(null)
  const test = async () => {
    await fetch(`${import.meta.env.VITE_BASE_API}/goodbye?name=test`, {
      method: 'PUT',
      headers: {
        'Content-type': 'application/json; charset=UTF-8',
      },
    })
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
      })
      .catch((err) => {
        console.log(err.message);
      });
  };
  useEffect(() => {
    console.log(wallets);
  });
  return (
    <>
      <div className="theme-dark bg-primary min-h-screen">
        <div className="flex mb-32">
          <div className="flex-1 w-14 h-14">
            <h1 className="text-3xl text-neutral">DNS MARKET</h1>
          </div>
          <div className="w-64 ...">
            <Wallets wallets={wallets} setWallets={setWallets} />
          </div>
        </div>

        <div className="grid grid-cols-2 justify-items-center">
          <div className="w-96 rounded-xl bg-secondary bg-clip-border text-gray-700 shadow-md">
            <p className="text-center mb-64 text-neutral">Register a new domain name</p>
            <Button onClick={() => {test()}}>Register</Button>
          </div>
          <div className="w-96 rounded-xl bg-secondary bg-clip-border text-gray-700 shadow-md">
            <p className="text-center mb-64 text-neutral">Your domain names</p>
          </div>
        </div>
      </div>
    </>
  )
}
