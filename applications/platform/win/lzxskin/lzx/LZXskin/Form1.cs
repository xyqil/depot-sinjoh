using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.IO;
using System.Linq;
using System.Text;
using System.Windows.Forms;
using System.Windows.Forms.VisualStyles;

namespace LZXskin
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void button3_Click(object sender, EventArgs e)//select file
        {
            OpenFileDialog openFileDialog1 = new OpenFileDialog();
            openFileDialog1.ShowDialog();
            FileInfo info = new FileInfo(openFileDialog1.FileName);
                string fileName = info.Name;
                string fileDir = info.DirectoryName;
                string fileInf = Path.Combine(fileDir, fileName);
                label1.Text = fileInf;
        }

        private void button1_Click(object sender, EventArgs e) // decompress
        {
            System.Diagnostics.Process process = new System.Diagnostics.Process();
            System.Diagnostics.ProcessStartInfo startInfo = new System.Diagnostics.ProcessStartInfo();
            startInfo.WindowStyle = System.Diagnostics.ProcessWindowStyle.Hidden;
            startInfo.FileName = "cmd.exe";
            startInfo.Arguments = "/k" + label2.Text + " -d " + label1.Text;
            process.StartInfo = startInfo;
            if (label1.Text == "No file selected.")
                MessageBox.Show("No file selected.");
            else if (label2.Text == "lzx location is not set.")
                MessageBox.Show("lzx location is not set.");
            else
            {
                process.Start();
                label3.ForeColor = System.Drawing.Color.Blue;
                label3.Text = "File decompressed.";
            }
        }

        private void button4_Click(object sender, EventArgs e) //select lzx
        {
            OpenFileDialog openFileDialog1 = new OpenFileDialog();
            openFileDialog1.ShowDialog();
            FileInfo info = new FileInfo(openFileDialog1.FileName);
            string lzxFileName = info.Name;
            string lzxFileDir = info.DirectoryName;
            string lzx = Path.Combine(lzxFileDir, lzxFileName);
            label2.Text = lzx;
            string path = "lzxskin.cfg";
            File.WriteAllText(path, lzx);
        }

        private void button2_Click(object sender, EventArgs e) //compress
        {
            System.Diagnostics.Process process = new System.Diagnostics.Process();
            System.Diagnostics.ProcessStartInfo startInfo = new System.Diagnostics.ProcessStartInfo();
            startInfo.WindowStyle = System.Diagnostics.ProcessWindowStyle.Hidden;
            startInfo.FileName = "cmd.exe";
            startInfo.Arguments = "/k" + label2.Text + " -evb " + label1.Text;
            process.StartInfo = startInfo;
            startInfo.RedirectStandardOutput = true;
            startInfo.UseShellExecute = false;
            if (label1.Text == "No file selected.")
                MessageBox.Show("No file selected.");
            else if (label2.Text == "lzx location is not set.")
                MessageBox.Show("lzx location is not set.");
            else
            {
                process.Start();
                label3.ForeColor = System.Drawing.Color.Red;
                label3.Text = "File compressed.";
            }
        }

        private void Form1_Load(object sender, EventArgs e)
        {
            string cfg = "lzxskin.cfg";
            if (File.Exists(cfg))
                label2.Text = File.ReadAllText(cfg);
        }
    }
}