console.log("Start...\n");

async function throwErrorImmediately(errorMessage) {
    throw new Error(errorMessage);
}

async function throwErrorAfterAwait(errorMessage) {
    await Promise.resolve();
    throw new Error(errorMessage);
}

throwErrorImmediately("Error #1").catch((error) => {
    console.error("Caught error:", error.message);
});

throwErrorAfterAwait("Error #2").catch((error) => {
    console.error("Caught error:", error.message);
});

try {
    // will brake the app
    throwErrorAfterAwait("Error #3");
} catch(error) {
    // This won't be printed because the error is thrown asynchronously.
    console.error("Caught error:", error.message);
}

console.log("\nEnd...\n");
