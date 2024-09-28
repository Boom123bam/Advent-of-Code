public class Amp {

    private int programIdx;
    private int[] program;
    public int phase;

    public Amp(int[] program, int phase) {
        this.programIdx = 0;
        this.program = program;
        this.phase = phase;
    }

    public int run(int[] inputs, boolean part2) {
        int lastOutput = -1;
        int inputIdx = 0;
        while (true) {
            int instruction = this.program[this.programIdx];
            int opcode = instruction % 100;
            int[] modes = { 0, 0, 0 };
            instruction /= 100;
            for (int j = 0; j < 3; j++) {
                modes[j] = instruction % 10;
                instruction /= 10;
            }

            int i1, i2, i3;
            i1 = this.programIdx + 1 < this.program.length
                ? this.program[this.programIdx + 1]
                : -1;
            i2 = this.programIdx + 2 < this.program.length
                ? this.program[this.programIdx + 2]
                : -1;
            i3 = this.programIdx + 3 < this.program.length
                ? this.program[this.programIdx + 3]
                : -1;
            switch (opcode) {
                case 1:
                    assert modes[2] == 0;
                    this.program[i3] = (modes[0] == 1 ? i1 : this.program[i1]) +
                    (modes[1] == 1 ? i2 : this.program[i2]);
                    this.programIdx += 4;
                    break;
                case 2:
                    assert modes[2] == 0;
                    this.program[i3] = (modes[0] == 1 ? i1 : this.program[i1]) *
                    (modes[1] == 1 ? i2 : this.program[i2]);
                    this.programIdx += 4;
                    break;
                case 3:
                    this.program[i1] = inputs[inputIdx];
                    // System.out.print(inputIdx);
                    inputIdx++;
                    this.programIdx += 2;
                    break;
                case 4:
                    if (part2) {
                        this.programIdx += 2;
                        return this.program[i1];
                    }
                    lastOutput = this.program[i1];
                    this.programIdx += 2;
                    break;
                case 5:
                    if ((modes[0] == 1 ? i1 : this.program[i1]) != 0) {
                        this.programIdx = modes[1] == 1 ? i2 : this.program[i2];
                    } else {
                        this.programIdx += 3;
                    }
                    break;
                case 6:
                    if ((modes[0] == 1 ? i1 : this.program[i1]) == 0) {
                        this.programIdx = modes[1] == 1 ? i2 : this.program[i2];
                    } else {
                        this.programIdx += 3;
                    }
                    break;
                case 7:
                    if (
                        (modes[0] == 1 ? i1 : this.program[i1]) <
                        (modes[1] == 1 ? i2 : this.program[i2])
                    ) {
                        this.program[i3] = 1;
                    } else {
                        this.program[i3] = 0;
                    }
                    this.programIdx += 4;
                    break;
                case 8:
                    if (
                        (modes[0] == 1 ? i1 : this.program[i1]) ==
                        (modes[1] == 1 ? i2 : this.program[i2])
                    ) {
                        this.program[i3] = 1;
                    } else {
                        this.program[i3] = 0;
                    }
                    this.programIdx += 4;
                    break;
                default:
                    assert opcode == 99;
                    // System.out.println("halt, " + lastOutput);
                    return lastOutput;
            }
        }
    }
}
