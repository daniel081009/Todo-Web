import Head from "next/head";

export default function Heads(title : string,content : string) {
    return (
        <Head>
            <title>{title}</title>
            <meta name="description" content={content} />
            <link rel="icon" href="/favicon.ico" />
        </Head>
    )
}