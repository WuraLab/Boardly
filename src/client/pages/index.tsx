import Head from 'next/head'
import Link from 'next/link'
import { GetStaticProps } from 'next'

export default function Home() {
  return (
    <div>
      <Head>
        <title>Boardly</title>
      </Head>
      <div>
         Welcome to Boardly
      </div>
    </div>
  )
}
