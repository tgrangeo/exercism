export class Matrix {
    matrix: number[][] = [];

    constructor(str: string) {
        let rows = str.split("\n");
        for (let i = 0; i < rows.length; i++) {
            if (!this.matrix[i]) this.matrix[i] = []; 
            let numbers = rows[i].split(" ").map(Number);
            this.matrix[i].push(...numbers);
        }
    }

    get rows(): number[][] {
        return this.matrix;
    }

    get columns(): number[][] {
        let cols: number[][] = [];
        for (let i = 0; i < this.matrix.length; i++) {
            for (let j = 0; j < this.matrix[i].length; j++) {
                if (cols[j] === undefined) {
                    cols[j] = []; 
                }
                cols[j][i] = this.matrix[i][j]; 
            }
        }
        return cols;
    }
}
