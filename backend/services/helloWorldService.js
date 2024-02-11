const getHelloWorldMessage = () => {
    return `hello world from service. testing env: ${ process.env.NEXT_PUBLIC_TEST_KEY }`;
};

module.exports = {
    getHelloWorldMessage
};