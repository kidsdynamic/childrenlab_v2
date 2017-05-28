export class Dashboard {
  totalUserCount: number;
  signup: DashboardSignup[];
  totalActivityCount: number;
  activity: DashboardActivity[];
  activityByEventDate: DashboardActivity[];
}

class DashboardSignup {
  signupCount: number;
  date: string;
}
class DashboardActivity {
  activityCount: number;
  userCount: number;
  indoorActivity: number;
  outdoorActivity: number;
  date: string;
}
