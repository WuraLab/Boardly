import Head from 'next/head'

// import Link from 'next/link';
// import { GetStaticProps } from 'next';
// import { Layout } from '../components/layout/index';

const Home: React.FC = () => {
    return (
        <>
            {/* <Layout> */}
            <div>
                <Head>
                    <title>Boardly</title>
                </Head>
                <div>Welcome to Boardly</div>
            </div>
            {/* </Layout> */}
        </>
    );
};

export default Home;
