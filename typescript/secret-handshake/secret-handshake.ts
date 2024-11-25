const action = ["wink", "double blink", "close your eyes", "jump", "reverse"];

export function commands(n: number): string[] {
    if (n === 0) return [];
    const binary = n.toString(2).padStart(5, "0");
    let ret: string[] = [];
    for (let i = 0; i < 5; i++) {
        if (binary[binary.length - 1 - i] === "1") {
            ret.push(action[i]);
        }
    }
    if (ret.includes("reverse")) {
        ret = ret.filter((act) => act !== "reverse");
        ret.reverse();
    }

    return ret;
}
