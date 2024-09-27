import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.*;

public class Day05 {

    public static void main(String[] args) {
        String line = "";
        String filePath = "input.txt";
        try (BufferedReader br = new BufferedReader(new FileReader(filePath))) {
            line = br.readLine();
        } catch (IOException e) {
            e.printStackTrace();
        }

        int[] originalArray = Arrays.stream(line.split(","))
            .mapToInt(Integer::parseInt)
            .toArray();

        int[] newArray = originalArray.clone();
        System.out.println(run(newArray, 1));
        newArray = originalArray.clone();
        System.out.println(run(newArray, 5));
    }

    public static int run(int[] P, int INPUT) {
        int lastOutput = -1;
        int ip = 0;
        while (true) {
            int instruction = P[ip];
            int opcode = instruction % 100;
            int[] modes = { 0, 0, 0 };
            instruction /= 100;
            for (int j = 0; j < 3; j++) {
                modes[j] = instruction % 10;
                instruction /= 10;
            }

            int i1, i2, i3;
            i1 = ip + 1 < P.length ? P[ip + 1] : -1;
            i2 = ip + 2 < P.length ? P[ip + 2] : -1;
            i3 = ip + 3 < P.length ? P[ip + 3] : -1;
            switch (opcode) {
                case 1:
                    assert modes[2] == 0;
                    P[i3] = (modes[0] == 1 ? i1 : P[i1]) +
                    (modes[1] == 1 ? i2 : P[i2]);
                    ip += 4;
                    break;
                case 2:
                    assert modes[2] == 0;
                    P[i3] = (modes[0] == 1 ? i1 : P[i1]) *
                    (modes[1] == 1 ? i2 : P[i2]);
                    ip += 4;
                    break;
                case 3:
                    P[i1] = INPUT;
                    ip += 2;
                    break;
                case 4:
                    lastOutput = P[i1];
                    ip += 2;
                    break;
                case 5:
                    if ((modes[0] == 1 ? i1 : P[i1]) != 0) {
                        ip = modes[1] == 1 ? i2 : P[i2];
                    } else {
                        ip += 3;
                    }
                    break;
                case 6:
                    if ((modes[0] == 1 ? i1 : P[i1]) == 0) {
                        ip = modes[1] == 1 ? i2 : P[i2];
                    } else {
                        ip += 3;
                    }
                    break;
                case 7:
                    if (
                        (modes[0] == 1 ? i1 : P[i1]) <
                        (modes[1] == 1 ? i2 : P[i2])
                    ) {
                        P[i3] = 1;
                    } else {
                        P[i3] = 0;
                    }
                    ip += 4;
                    break;
                case 8:
                    if (
                        (modes[0] == 1 ? i1 : P[i1]) ==
                        (modes[1] == 1 ? i2 : P[i2])
                    ) {
                        P[i3] = 1;
                    } else {
                        P[i3] = 0;
                    }
                    ip += 4;
                    break;
                default:
                    assert opcode == 99;
                    return lastOutput;
            }
        }
    }
}
