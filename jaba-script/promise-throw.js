console.log("Start...\n");

new Promise(() => {
    console.log("Promise #1");
    // Promise automatically catches errors thrown in the executor and passes them to the `reject` callback
    throw new Error("Error #1");
}).catch((error) => {
    console.error("Caught error:", error.message);
});

try {
    new Promise(() => {
        console.log("Promise #2");
        // Even though promises are asynchronous, uncaught errors will still be able to brake the app
        throw new Error("Error #2");
    });
} catch(error) {
    // This won't be printed because the error is thrown asynchronously.
    console.error("Caught error:", error.message);
}

new Promise(() => {
    console.log("Promise #3");
});

console.log("\nEnd...\n");
