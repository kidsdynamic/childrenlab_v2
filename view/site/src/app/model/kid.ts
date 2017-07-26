/**
 * Created by yen-chiehchen on 4/16/17.
 */
import { User } from './user'

export class Kid {
  id: number;
  name: string;
  profile: string;
  macId: string;
  dateCreated: string;
  parent: User;

}

export class BatteryStatus {
  macId: string;
  batteryLife: number;
  dateReceived: string;
}
