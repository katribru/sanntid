with Ada.Text_IO, Ada.Integer_Text_IO, Ada.Numerics.Float_Random;
use Ada.Text_IO, Ada.Integer_Text_IO, Ada.Numerics.Float_Random;

procedure exercise7 is

	Count_Failed : exception; -- Exception to be raised when counting fails
	Gen : Generator; -- Random number generator
	protected type Transaction_Manager (N : Positive) is
		entry Finished;
		function Commit return Boolean;
		procedure Signal_Abort;
	private
		Finished_Gate_Open : Boolean := False;
		Aborted : Boolean := False;
		Should_Commit : Boolean := True;
	end Transaction_Manager;
	
	protected body Transaction_Manager is
		entry Finished when Finished_Gate_Open or Finished'Count = N is
		begin
			if  Finished'Count = (N-1) then
				Finished_Gate_Open := True;
			end if;
			
			if Aborted = True and Finished'Count = (N-1) then
				Should_Commit := False;
			end if;
			
			if Aborted = False and Finished'Count = (N-1) then
				Should_Commit := True; 
			end if;
			
			if Finished'Count = 0 then
				Finished_Gate_Open := False;
				Aborted := False; 
			end if;	
		end Finished;
		procedure Signal_Abort is
		begin
			Aborted := True;
		end Signal_Abort;
		function Commit return Boolean is
		begin
			return Should_Commit;
		end Commit;
	end Transaction_Manager;




	function Unreliable_Slow_Add (x : Integer) return Integer is
		Error_Rate : Constant := 0.15; -- (between 0 and 1)
		Rand_num : float := Random(Gen);
		d : float := 4.0*Random(Gen);
	begin

		if Rand_num < Error_Rate then
			delay Duration(0.5);
			raise Count_Failed;
		else
			delay Duration(d);
			return x + 10;
		end if;	
	end Unreliable_Slow_Add;



	task type Transaction_Worker (Initial : Integer; Manager : access Transaction_Manager);
	task body Transaction_Worker is
		Num : Integer := Initial;
		Prev : Integer := Num;
		Round_Num : Integer := 0;
	begin
		Put_Line ("Worker" & Integer'Image(Initial) & " started");
		loop
			Put_Line ("Worker" & Integer'Image(Initial) & " started round" & Integer'Image(Round_Num));
			Round_Num := Round_Num + 1;
			
			begin
				Num := Unreliable_Slow_Add(Num);
			exception
				when Count_Failed =>
					Manager.Signal_Abort;
			end; --Hum?
			Manager.Finished;		
			if Manager.Commit = True then
				Put_Line (" Worker" & Integer'Image(Initial) & " comitting" & Integer'Image(Num));
			else
				Put_Line (" Worker" & Integer'Image(Initial) &
					" reverting from" & Integer'Image(Num) &
					" to" & Integer'Image(Prev));
				Num := Prev; --Roll back	
			end if;

			Prev := Num;
			delay 0.5;
			
		end loop;
	end Transaction_Worker;

	Manager : aliased Transaction_Manager (3);

	Worker_1 : Transaction_Worker (0, Manager'Access);
	Worker_2 : Transaction_Worker (1, Manager'Access);
	Worker_3 : Transaction_Worker (2, Manager'Access);
begin
	Reset(Gen); -- Seed the random number generator
end exercise7;
