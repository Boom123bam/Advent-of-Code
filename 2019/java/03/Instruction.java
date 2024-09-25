public class Instruction {

    public int mag;
    public char dir;

    public Instruction(String instructionString) {
        this.dir = instructionString.charAt(0);
        this.mag = Integer.parseInt(instructionString.substring(1));
    }

    public int[] getDirVector() {
        switch (this.dir) {
            case 'U':
                return new int[] { 0, 1 };
            case 'D':
                return new int[] { 0, -1 };
            case 'L':
                return new int[] { -1, 0 };
            case 'R':
                return new int[] { 1, 0 };
            default:
                throw new IllegalArgumentException(
                    "Invalid direction: " + this.dir
                );
        }
    }

    @Override
    public String toString() {
        return "" + this.dir + this.mag;
    }
}
