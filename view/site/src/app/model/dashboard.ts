export class Dashboard {
  totalUserCount: number;
  signup: DashboardSignup[];
  totalActivityCount: number;
  activity: DashboardActivity[];
}

class DashboardSignup {
  signupCount: number;
  date: string;
}
class DashboardActivity {
  activityCount: number;
  userCount: number;
  date: string;
}
