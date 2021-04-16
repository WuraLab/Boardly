// // types
// type User = {
//     user_id: string;
//     email: string;
//     user_name: string;
//     password: string;
// };

// // anonymous function
// async function getUsers() {
//     // curl ""
//     const res = await fetch(`${API_URL}/readUsers`)
//         //convert and return json
//         .then((res) => res.json());

//     // console.log(res)
//     return res.users; // type object[]
// }

// // Nextjs will fetch data at build time and cached
// // getStaticProps // Data fetching feature
// export const getStaticProps = async ({ params }) => {
//     const users = await getUsers();

//     return {
//         revalidate: 10, // revalidate this data , from the backend every 10 seconds
//         props: { users } //syntax props:{}
//         // but Users is actually an array
//     };
// };
